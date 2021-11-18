package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"

	"github.com/ozonmp/srv-verification-api/internal/config"
)

func createStatusServer(cfg *config.Config, isReady *atomic.Value) *http.Server {
	statusAddr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)

	mux := http.DefaultServeMux

	mux.HandleFunc(cfg.Status.LivenessPath, livenessHandler)
	mux.HandleFunc(cfg.Status.ReadinessPath, readinessHandler(isReady))
	mux.HandleFunc(cfg.Status.VersionPath, versionHandler(cfg))

	statusServer := &http.Server{
		Addr:    statusAddr,
		Handler: mux,
	}

	return statusServer
}

func livenessHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)

			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func versionHandler(cfg *config.Config) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		data := map[string]interface{}{
			"name":        cfg.Project.Name,
			"debug":       cfg.Project.Debug,
			"environment": cfg.Project.Environment,
			"version":     cfg.Project.Version,
			"commitHash":  cfg.Project.CommitHash,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			logger.ErrorKV(context.Background(), "Service information encoding error", "err", err)
		}
	}
}

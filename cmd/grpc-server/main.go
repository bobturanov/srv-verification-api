package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/ozonmp/srv-verification-api/internal/config"
	"github.com/ozonmp/srv-verification-api/internal/database"
	"github.com/ozonmp/srv-verification-api/internal/server"
	"github.com/ozonmp/srv-verification-api/internal/tracer"
)

var (
	batchSize uint = 2
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, "Failed init configuration", "err", err)
	}
	cfg := config.GetConfigInstance()

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	syncLogger := logger.InitLogger(ctx, &cfg)
	defer syncLogger()

	logger.InfoKV(ctx, fmt.Sprintf("Starting service: %s", cfg.Project.Name),
		"version", cfg.Project.Version,
		"commitHash", cfg.Project.CommitHash,
		"debug", cfg.Project.Debug,
		"environment", cfg.Project.Environment,
	)

	initCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode)

	db, err := database.NewPostgres(initCtx, dsn, cfg.Database.Driver)
	if err != nil {
		logger.FatalKV(ctx, "Failed init postgres", "err", err)
	}
	defer db.Close()

	if *migration {
		if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			logger.ErrorKV(ctx, "Migration failed", "err", err)
			return
		}
	}

	tracing, err := tracer.NewTracer(&cfg)
	if err != nil {
		logger.ErrorKV(ctx, "Failed init tracing", "err", err)
		return
	}
	defer tracing.Close()

	if err := server.NewGrpcServer(db, batchSize).Start(&cfg); err != nil {
		logger.ErrorKV(ctx, "Failed creating gRPC server", "err", err)

		return
	}
}

package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/ozonmp/srv-verification-api/internal/repo"

	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
)

var (
	totalVerificationNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "srv_verification_api_verification_not_found_total",
		Help: "Total number of verifications that were not found",
	})
)

type verificationAPI struct {
	pb.UnimplementedSrvVerificationApiServiceServer
	repo repo.Repo
}

// NewVerificationAPI returns api of srv-verification-api service
func NewVerificationAPI(r repo.Repo) pb.SrvVerificationApiServiceServer {
	return &verificationAPI{repo: r}
}

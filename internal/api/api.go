package api

import (
	"context"
	"github.com/ozonmp/srv-verification-api/internal/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (o *verificationAPI) DescribeVerificationV1(
	ctx context.Context,
	req *pb.DescribeVerificationV1Request,
) (*pb.DescribeVerificationV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeVerificationV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	verification, err := o.repo.DescribeVerification(ctx, req.VerificationId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeVerificationV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if verification == nil {
		log.Debug().Uint64("verificationId", req.VerificationId).Msg("verification not found")
		totalVerificationNotFound.Inc()

		return nil, status.Error(codes.NotFound, "verification not found")
	}

	log.Debug().Msg("DescribeVerificationV1 - success")

	return &pb.DescribeVerificationV1Response{
		Value: &pb.Verification{
			Id:  verification.ID,
			Name: verification.Name,
		},
	}, nil
}

func (o *verificationAPI) CreateVerificationV1(
	ctx context.Context,
	req *pb.CreateVerificationV1Request,
) (*pb.CreateVerificationV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateVerificationV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	verification := model.Verification{Name: req.VerificationName}
	if err := o.repo.AddVerification(ctx, &verification); err != nil {
		log.Error().Err(err).Msg("DescribeVerificationV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateVerificationV1 - success")

	return &pb.CreateVerificationV1Response{
		VerificationId: verification.ID,
	}, nil
}

func (o *verificationAPI) ListVerificationV1(
	ctx context.Context,
	req *pb.ListVerificationV1Request,
) (*pb.ListVerificationV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListVerificationV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	verifications, err := o.repo.ListVerification(ctx)
	if err != nil {
		log.Error().Err(err).Msg("ListVerificationV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("ListVerificationV1 - success")

	verificationsPb := make([]*pb.Verification, len(verifications))

	for _, verification := range verifications {
		verificationsPb = append(verificationsPb, convertVerificationToPb(verification))
	}

	return  &pb.ListVerificationV1Response{
		Verification: verificationsPb,
	}, nil
}

func convertVerificationToPb(verification *model.Verification) *pb.Verification {
	return &pb.Verification{
		Id:   verification.ID,
		Name: verification.Name,
	}
}

func (o *verificationAPI) RemoveVerificationV1(
	ctx context.Context,
	req *pb.RemoveVerificationV1Request,
) (*pb.RemoveVerificationV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveVerificationV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.repo.RemoveVerification(ctx, req.VerificationId)
	if err != nil {
		log.Error().Err(err).Msg("RemoveVerificationV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !result {
		log.Debug().Uint64("verificationId", req.VerificationId).Msg("verification not found")
		totalVerificationNotFound.Inc()

		return nil, status.Error(codes.NotFound, "verification not found")
	}

	log.Debug().Msg("DescribeVerificationV1 - success")

	return &pb.RemoveVerificationV1Response{
		Result: result,
	}, nil
}

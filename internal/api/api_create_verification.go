package api

import (
	"context"

	"github.com/ozonmp/srv-verification-api/internal/model"
	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

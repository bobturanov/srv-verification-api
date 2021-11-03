package api

import (
	"context"

	"github.com/ozonmp/srv-verification-api/internal/model"
	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

	return &pb.ListVerificationV1Response{
		Verification: verificationsPb,
	}, nil
}

func convertVerificationToPb(verification *model.Verification) *pb.Verification {
	return &pb.Verification{
		Id:   verification.ID,
		Name: verification.Name,
	}
}

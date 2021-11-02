package api

import (
	"context"
	"github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *verificationAPI) RemoveVerificationV1(
	ctx context.Context,
	req *srv_verification_api.RemoveVerificationV1Request,
) (*srv_verification_api.RemoveVerificationV1Response, error) {

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

	return &srv_verification_api.RemoveVerificationV1Response{
		Result: result,
	}, nil
}

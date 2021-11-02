package api

import (
	"context"
	"github.com/ozonmp/srv-verification-api/internal/model"
	"github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *verificationAPI) ListVerificationV1(
	ctx context.Context,
	req *srv_verification_api.ListVerificationV1Request,
) (*srv_verification_api.ListVerificationV1Response, error) {

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

	verificationsPb := make([]*srv_verification_api.Verification, len(verifications))

	for _, verification := range verifications {
		verificationsPb = append(verificationsPb, convertVerificationToPb(verification))
	}

	return &srv_verification_api.ListVerificationV1Response{
		Verification: verificationsPb,
	}, nil
}

func convertVerificationToPb(verification *model.Verification) *srv_verification_api.Verification {
	return &srv_verification_api.Verification{
		Id:   verification.ID,
		Name: verification.Name,
	}
}


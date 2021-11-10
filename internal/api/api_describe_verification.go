package api

import (
	"context"

	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
			Value: convertVerificationToPb(verification),
		},
		nil
}

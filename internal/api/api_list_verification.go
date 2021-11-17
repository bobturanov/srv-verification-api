package api

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"

	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *verificationAPI) ListVerificationV1(
	ctx context.Context,
	req *pb.ListVerificationV1Request,
) (*pb.ListVerificationV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.ListVerificationV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListVerificationV1 - invalid argument", "err", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	verifications, err := o.repo.ListVerification(ctx)
	if err != nil {
		logger.ErrorKV(ctx, "ListVerificationV1 -- failed", "err", err)

		return nil, status.Error(codes.Internal, err.Error())
	}
	logger.DebugKV(ctx, "ListVerificationV1 - success")

	verificationsPb := make([]*pb.Verification, 0, len(verifications))

	for _, verification := range verifications {
		verificationsPb = append(verificationsPb, convertVerificationToPb(verification))
	}

	return &pb.ListVerificationV1Response{
		Verification: verificationsPb,
	}, nil
}

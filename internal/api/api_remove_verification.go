package api

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"

	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *verificationAPI) RemoveVerificationV1(
	ctx context.Context,
	req *pb.RemoveVerificationV1Request,
) (*pb.RemoveVerificationV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.RemoveVerificationV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "RemoveVerificationV1 - invalid argument", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.repo.RemoveVerification(ctx, req.VerificationId)
	if err != nil {
		logger.ErrorKV(ctx, "RemoveVerificationV1 -- failed", "err", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !result {
		logger.DebugKV(ctx, fmt.Sprintf("verification not found - verificationId: %d", req.VerificationId))
		totalVerificationNotFound.Inc()

		return nil, status.Error(codes.NotFound, "verification not found")
	}
	logger.DebugKV(ctx, "DescribeVerificationV1 - success")

	return &pb.RemoveVerificationV1Response{
		Result: result,
	}, nil
}

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

func (o *verificationAPI) DescribeVerificationV1(
	ctx context.Context,
	req *pb.DescribeVerificationV1Request,
) (*pb.DescribeVerificationV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.DescribeVerificationV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "DescribeVerificationV1 - invalid argument", "err", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	verification, err := o.repo.DescribeVerification(ctx, req.VerificationId)
	if err != nil {
		logger.ErrorKV(ctx, "DescribeVerificationV1 -- failed", "err", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if verification == nil {
		logger.DebugKV(ctx, fmt.Sprintf("verification not found - verificationId: %d", req.VerificationId))
		totalVerificationNotFound.Inc()

		return nil, status.Error(codes.NotFound, "verification not found")
	}
	logger.DebugKV(ctx, "DescribeVerificationV1 - success")

	return &pb.DescribeVerificationV1Response{
			Value: convertVerificationToPb(verification),
		},
		nil
}

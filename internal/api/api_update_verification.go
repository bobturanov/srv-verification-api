package api

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"

	"github.com/ozonmp/srv-verification-api/internal/model"
	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *verificationAPI) UpdateVerificationV1(
	ctx context.Context,
	req *pb.UpdateVerificationV1Request,
) (*pb.UpdateVerificationV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.UpdateVerificationV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "UpdateVerificationV1 - invalid argument", "err", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	verification := model.Verification{ID: req.VerificationId, Name: req.Name, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	dbStatus, err := o.repo.UpdateVerification(ctx, &verification)
	if err != nil {
		logger.ErrorKV(ctx, "UpdateVerificationV1 -- failed", "err", err)

		return nil, status.Error(codes.Internal, err.Error())
	}
	logger.DebugKV(ctx, "UpdateVerificationV1 - success")

	return &pb.UpdateVerificationV1Response{
		Result: dbStatus,
	}, nil
}

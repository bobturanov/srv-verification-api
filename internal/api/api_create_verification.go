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

func (o *verificationAPI) CreateVerificationV1(
	ctx context.Context,
	req *pb.CreateVerificationV1Request,
) (*pb.CreateVerificationV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.CreateVerificationV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "CreateVerificationV1 - invalid argument", "err", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	verification := model.Verification{Name: req.VerificationName, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := o.repo.AddVerification(ctx, &verification); err != nil {
		logger.ErrorKV(ctx, "CreateVerificationV1 -- failed", "err", err)

		return nil, status.Error(codes.Internal, err.Error())
	}
	logger.DebugKV(ctx, "CreateVerificationV1 - success")

	return &pb.CreateVerificationV1Response{
		VerificationId: verification.ID,
	}, nil
}

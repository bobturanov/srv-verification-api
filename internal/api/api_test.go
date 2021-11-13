package api

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/srv-verification-api/internal/mocks"
	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var errNotImplementedMethod = errors.New("method is not implemented")

type VerificationAPITestSuite struct {
	suite.Suite
	listener *bufconn.Listener
	server   *grpc.Server
	conn     *grpc.ClientConn
	client   pb.SrvVerificationApiServiceClient
}

func (s *VerificationAPITestSuite) bufDialer(context.Context, string) (net.Conn, error) {
	return s.listener.Dial()
}

//nolint
func (s *VerificationAPITestSuite) SetupSuite() {
	s.listener = bufconn.Listen(bufSize)
	s.server = grpc.NewServer()

	ctrl := gomock.NewController(s.T())
	repo := mocks.NewMockRepo(ctrl)
	repo.EXPECT().AddVerification(gomock.Any(), gomock.Any()).Return(errNotImplementedMethod)
	repo.EXPECT().DescribeVerification(gomock.Any(), gomock.Any()).Return(nil, errNotImplementedMethod)
	repo.EXPECT().ListVerification(gomock.Any()).Return(nil, errNotImplementedMethod)
	repo.EXPECT().RemoveVerification(gomock.Any(), gomock.Any()).Return(false, errNotImplementedMethod)

	pb.RegisterSrvVerificationApiServiceServer(s.server, NewVerificationAPI(repo))
	go func() {
		if err := s.server.Serve(s.listener); err != nil {
			log.Fatalf("s.server exited with error: %v", err)
		}
	}()

	ctx := context.Background()
	var err error
	s.conn, err = grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(s.bufDialer), grpc.WithInsecure())
	if err != nil {
		s.T().Fatalf("failed to dial bufnet: %v", err)
	}

	s.client = pb.NewSrvVerificationApiServiceClient(s.conn)

}

func (s *VerificationAPITestSuite) DownSuite() {
	err := s.conn.Close()
	if err != nil {
		log.Panic(err)
	}
	s.server.Stop()
}

//nolint
func (s *VerificationAPITestSuite) TestCreateVerification() {
	req := &pb.CreateVerificationV1Request{
		VerificationName: "TestName",
	}
	resp, err := s.client.CreateVerificationV1(context.Background(), req)
	s.Nil(resp)
	s.NotNil(err)

	st, _ := status.FromError(err)
	s.Equal(codes.Internal, st.Code())
	s.Equal("method is not implemented", st.Message())

}

//nolint
func (s *VerificationAPITestSuite) TestDescribeVerification() {
	req := &pb.DescribeVerificationV1Request{
		VerificationId: 2236,
	}
	resp, err := s.client.DescribeVerificationV1(context.Background(), req)
	s.Nil(resp)
	s.NotNil(err)

	st, _ := status.FromError(err)
	s.Equal(codes.Internal, st.Code())
	s.Equal("method is not implemented", st.Message())

}

//nolint
func (s *VerificationAPITestSuite) TestListeVerification() {
	req := &pb.ListVerificationV1Request{}
	resp, err := s.client.ListVerificationV1(context.Background(), req)
	s.Nil(resp)
	s.NotNil(err)

	st, _ := status.FromError(err)
	s.Equal(codes.Internal, st.Code())
	s.Equal("method is not implemented", st.Message())

}

//nolint
func (s *VerificationAPITestSuite) TestRemoveVerification() {
	req := &pb.RemoveVerificationV1Request{VerificationId: 754}
	resp, err := s.client.RemoveVerificationV1(context.Background(), req)
	s.Nil(resp)
	s.NotNil(err)

	st, _ := status.FromError(err)
	s.Equal(codes.Internal, st.Code())
	s.Equal("method is not implemented", st.Message())

}

func TestVerificationAPI(t *testing.T) {
	suite.Run(t, new(VerificationAPITestSuite))
}
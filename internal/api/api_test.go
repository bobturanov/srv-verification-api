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

func (s *VerificationAPITestSuite) SetupSuite() {
	s.listener = bufconn.Listen(bufSize)
	s.server = grpc.NewServer()

	ctrl := gomock.NewController(s.T())
	repo := mocks.NewMockRepo(ctrl)
	repo.EXPECT().AddVerification(gomock.Any(), gomock.Any()).Return(errors.New("method is not implemented"))

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

func TestLocationAPI(t *testing.T) {
	suite.Run(t, new(VerificationAPITestSuite))
}

//func initAPI(t *testing.T) {
//	lis = bufconn.Listen(bufSize)
//	s := grpc.NewServer()
//	ctrl := gomock.NewController(t)
//	repo := mocks.NewMockRepo(ctrl)
//
//	pb.RegisterSrvVerificationApiServiceServer(s, NewVerificationAPI(repo))
//	go func() {
//		if err := s.Serve(lis); err != nil {
//			log.Fatalf("Server exited with error: %v", err)
//		}
//	}()
//}

//func TestSayHello(t *testing.T) {
//	initAPI(t)
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("Failed to dial bufnet: %v", err)
//	}
//	defer conn.Close()
//	client := pb.NewSrvVerificationApiServiceClient(conn)
//	resp, err := client.CreateVerificationV1(context.Background(), &pb.CreateVerificationV1Request{VerificationName: "Test"}) //TODO не те методы или не того интерфейса мокаются вроде (?)
//	if err != nil {
//		t.Fatalf("SayHello failed: %v", err)
//	}
//	log.Printf("Response: %+v", resp)
//	// Test for output here.
//}

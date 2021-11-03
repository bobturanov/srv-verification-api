package api

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/srv-verification-api/internal/mocks"
	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func initAPI(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockRepo(ctrl)

	pb.RegisterSrvVerificationApiServiceServer(s, NewVerificationAPI(repo))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSayHello(t *testing.T) {
	initAPI(t)
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewSrvVerificationApiServiceClient(conn)
	resp, err := client.CreateVerificationV1(ctx, &pb.CreateVerificationV1Request{VerificationName: "Test"}) //TODO не те методы или не того интерфейса мокаются вроде (?)
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}

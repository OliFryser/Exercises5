package server

import (
	"context"

	"google.golang.org/grpc"
)

type Server struct {
	// an interface that the server needs to have
	gRPC.UnimplementedTemplateServer

	// here you can impliment other fields that you want
}

func main() {
	grpcServer := grpc.NewServer(opts...)
	server := &Server
	gRPC.RegisterTemplateServer(grpcServer, server)
	grpcServer.Serve(list)
}

func (s *Server) GetServerTime(ctx context.Context, t1 *TimeRequest) (*TimeResponse, error) {
	t2 := "time.Now()"
	tr := TimeResponse{
		t2: "bob",
		t3: "lop",
	}
	timeResponse := new(TimeResponse())

	//return (tr, nil)
}

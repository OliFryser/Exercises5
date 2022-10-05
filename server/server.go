package main

import (
	"context"
	"net"

	"github.com/OliFryser/Exercises5/proto"
	"google.golang.org/grpc"
)

type Server struct {
	// an interface that the server needs to have
	proto.UnimplementedTimeServiceServer

	// here you can impliment other fields that you want
}

func main() {
	list, _ := net.Listen("tcp", ":9080")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := &Server{}
	proto.RegisterTimeServiceServer(grpcServer, server)
	grpcServer.Serve(list)
}

func (s *Server) GetServerTime(ctx context.Context, t1 *proto.TimeRequest) (*proto.TimeResponse, error) {
	t2 := "time.Now()"
	tr := proto.TimeResponse{
		T2: t2,
		T3: "lop",
	}
	return &tr, nil
}

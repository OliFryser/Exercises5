package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/OliFryser/Exercises5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Same principle as in client. Flags allows for user specific arguments/values
var clientsName = flag.String("name", "default", "Senders name")
var serverPort = flag.String("server", "5400", "Tcp server")

var server proto.TimeServiceClient //the server
var ServerConn *grpc.ClientConn    //the server connection

func main() {
	var opts []grpc.DialOption
	opts = append(
		opts, grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	timeContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, _ := grpc.DialContext(
		timeContext,
		fmt.Sprintf(":%s", *serverPort),
		opts...,
	)

	server = proto.NewTimeServiceClient(conn)

	treq := proto.TimeRequest{
		T1: "Current Time",
	}
	server.GetServerTime(timeContext, &treq)

}

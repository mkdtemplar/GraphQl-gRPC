package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var add = "0.0.0.0:50051"
var ConnectClient *grpc.ClientConn

func clientConn() *grpc.ClientConn {
	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial %v\n", err)
	}

	ConnectClient = conn

	return conn
}

func GetClientConnection() *grpc.ClientConn {
	return ConnectClient
}

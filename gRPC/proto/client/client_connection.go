package client

import (
	"graphqhhowto/gRPC/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func clientConn() (*grpc.ClientConn, proto.UserServiceClient, error) {
	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {

		}
	}(conn)

	c := proto.NewUserServiceClient(conn)

	return conn, c, nil
}

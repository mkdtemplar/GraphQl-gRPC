package client

import (
	"context"
	"graphqhhowto/database"
	"graphqhhowto/gRPC/proto"
	pb "graphqhhowto/gRPC/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var add = ":50051"

func CreateUserInDb(in *database.User) *proto.User {
	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {

		}
	}(conn)

	c := pb.NewUserServiceClient(conn)
	r, err := c.CreateUser(context.Background(), &proto.User{
		Id:   in.ID.String(),
		Name: in.Name,
	})
	return r
}

package client

import (
	"context"
	pb "graphqhhowto/gRPC/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListAllUsers() []*pb.User {
	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {

		}
	}(conn)
	var users []*pb.User
	c := pb.NewUserServiceClient(conn)

	stream, err := c.GetAllUsers(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Cannot list users from database %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receivind stream %v\n", err)
		}
		users = append(users, res)
	}
	return users
}

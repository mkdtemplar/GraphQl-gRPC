package client

import (
	"context"
	pb "graphqhhowto/gRPC/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	*grpc.ClientConn
}

func ListAllUsers() []*pb.User {
	var users []*pb.User
	conn, c, err := clientConn()
	if err != nil {
		return nil
	}

	c = pb.NewUserServiceClient(conn)

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

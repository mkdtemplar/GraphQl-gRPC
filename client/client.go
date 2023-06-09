package client

import (
	"context"
	pb "gRPC-services/gRPC/proto"
	"graphqhhowto/database"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	*grpc.ClientConn
}

type UserInterfaces interface {
	CreateUserInDb(in *database.User) *pb.User
	ListAllUsers() []*pb.User
}

func NewClient() UserInterfaces {
	return &Client{GetClientConnection()}
}

var Cl = NewClient()

func (c *Client) CreateUserInDb(in *database.User) *pb.User {
	c.ClientConn = clientConn()

	client := pb.NewUserServiceClient(c.ClientConn)

	r, err := client.CreateUser(context.Background(), &pb.User{
		Id:   in.ID.String(),
		Name: in.Name,
	})
	if err != nil {
		return nil
	}
	return r
}

func (c *Client) ListAllUsers() []*pb.User {
	var users []*pb.User
	c.ClientConn = clientConn()
	client := pb.NewUserServiceClient(c.ClientConn)

	stream, err := client.GetAllUsers(context.Background(), &emptypb.Empty{})
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

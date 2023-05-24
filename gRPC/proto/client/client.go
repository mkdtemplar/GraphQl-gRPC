package client

import (
	"context"
	"graphqhhowto/database"
	pb "graphqhhowto/gRPC/proto"

	"google.golang.org/grpc"
)

type Client struct {
	*grpc.ClientConn
}

type UserInterfaces interface {
	CreateUserInDb(in *database.User) *pb.User
}

func NewClient() UserInterfaces {
	return &Client{GetClientConnection()}
}

var Cl = NewClient()

func (c *Client) CreateUserInDb(in *database.User) *pb.User {
	c.ClientConn = clientConn()

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(c.ClientConn)
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

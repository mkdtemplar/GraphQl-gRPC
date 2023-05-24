package client

import (
	"context"
	"graphqhhowto/database"
	"graphqhhowto/gRPC/proto"
	pb "graphqhhowto/gRPC/proto"
)

var add = ":50051"

func CreateUserInDb(in *database.User) *proto.User {
	conn, c, err := clientConn()
	if err != nil {
		return nil
	}

	c = pb.NewUserServiceClient(conn)
	r, err := c.CreateUser(context.Background(), &proto.User{
		Id:   in.ID.String(),
		Name: in.Name,
	})
	return r
}

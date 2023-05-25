package main

import (
	"context"
	"fmt"
	"graphqhhowto/database"
	"graphqhhowto/gRPC/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, in *proto.User) (*proto.User, error) {
	id, _ := uuid.Parse(in.Id)
	data := &database.User{
		ID:   id,
		Name: in.Name,
	}

	res, err := data.Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error %v\n", err))
	}

	return &proto.User{
		Id:   res.ID.String(),
		Name: res.Name,
	}, nil
}

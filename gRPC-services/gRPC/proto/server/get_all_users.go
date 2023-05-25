package main

import (
	"fmt"
	"graphqhhowto/database"
	pb "graphqhhowto/gRPC/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetAllUsers(in *emptypb.Empty, stream pb.UserService_GetAllUsersServer) error {
	var user database.User
	usersDB, err := user.GetAllUsers()
	if err != nil {
		log.Println(err)
	}
	for _, u := range *usersDB {
		err = stream.Send(&pb.User{
			Id:   u.ID.String(),
			Name: u.Name,
		})
		if err != nil {
			return status.Errorf(codes.DataLoss, fmt.Sprintf("Error sending data %v\n", err))
		}
	}

	return nil
}

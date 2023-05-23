package main

import (
	"graphqhhowto/database"
	pb "graphqhhowto/gRPC/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var add = ":50051"

type Server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	database.InitDB()

	lis, err := net.Listen("tcp", add)

	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}

	log.Printf("Listening to address %s\n", add)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server %v\n", err)
	}
}

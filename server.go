package main

import (
	"context"
	"graphqhhowto/database"
	"graphqhhowto/graph"
	"log"
	"net"
	"net/http"
	"os"

	pb "graphqhhowto/gRPC/proto"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

type Server struct {
	pb.UserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, usr *pb.User) (*pb.User, error) {
	//userId, _ := uuid.Parse(usr.Id)
	//data := &database.User{
	//	ID:   userId,
	//	Name: usr.Name,
	//}
	//
	//res, _ := data.Save(ctx)

	return nil, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	database.InitDB()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}

	log.Printf("Listening to address %s\n", port)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server %v\n", err)
	}
	mux := http.NewServeMux()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

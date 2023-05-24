package main

import (
	"graphqhhowto/graph"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "0.0.0.0:8080"

func graphHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gin.SetMode(gin.ReleaseMode)
	mux := gin.Default()

	mux.POST("/query", graphHandler())
	mux.GET("/", playgroundHandler())
	err := mux.Run(defaultPort)
	if err != nil {
		log.Fatalf("Can not start server %s", err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}

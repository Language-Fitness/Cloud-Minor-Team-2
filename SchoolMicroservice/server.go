package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"school/graph"
	"school/internal/auth"
	"school/internal/service"
	"school/proto/pb"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8083"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenMiddleware := auth.Middleware

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))

	http.Handle("/query", tokenMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	//go grpcSchoolServer()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func grpcSchoolServer() {
	lis, err := net.Listen("tcp", ":9050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSchoolServiceServer(grpcServer, service.NewSchoolGRPCService())

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

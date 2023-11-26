package main

import (
	"ResultMicroservice/graph"
	"ResultMicroservice/grpcServers"
	"ResultMicroservice/proto/result"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	grpcServer()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func grpcServer() {
	// Replace with the actual path to your server certificate
	//creds, err := credentials.NewServerTLSFromFile("path/to/your/server.crt", "path/to/your/server.key")
	//if err != nil {
	//	log.Fatalf("Failed to load TLS keys: %v", err)
	//}
	//
	//// Create a new gRPC server with the TLS credentials
	//server := grpc.NewServer(grpc.Creds(creds))

	// Create a new gRPC server without TLS
	server := grpc.NewServer()

	// Register the server with the generated protobuf code
	result.RegisterGrpcResultServer(server, &grpcServers.ResultServer{})

	// Create a listener on a specific port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server is listening on port 50051...")

	// Serve the gRPC server
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

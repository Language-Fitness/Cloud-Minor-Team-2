package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"school/graph"
	"school/internal/auth"

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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//func grpcSagaServer() {
//	lis, err := net.Listen("tcp", ":9092")
//	if err != nil {
//		log.Fatalf("Failed to listen: %v", err)
//	}
//
//	grpcServer := grpc.NewServer()
//	pb.RegisterGRPCSagaServiceServer(grpcServer, service.NewSagaService())
//
//	log.Printf("server listening at %v", lis.Addr())
//	if err := grpcServer.Serve(lis); err != nil {
//		log.Fatalf("Failed to serve: %v", err)
//	}
//}

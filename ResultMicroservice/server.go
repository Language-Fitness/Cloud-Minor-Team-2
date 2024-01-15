package main

import (
	"ResultMicroservice/graph"
	"ResultMicroservice/internal/auth"
	"ResultMicroservice/internal/database"
	"ResultMicroservice/internal/database/migrations"
	service "ResultMicroservice/internal/rpc"
	"ResultMicroservice/proto/pb"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8085"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	migrations.InitResult()
	// Prometheus Metrics
	http.Handle("/metrics", promhttp.Handler())

	tokenMiddleware := auth.Middleware

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))

	http.Handle("/query", tokenMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	go grpcSagaServer()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func grpcSagaServer() {
	lis, err := net.Listen("tcp", ":9094")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	collection, _ := database.GetCollection()
	pb.RegisterGRPCSagaServiceServer(grpcServer, service.NewSagaService(collection))

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

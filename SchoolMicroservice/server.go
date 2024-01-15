package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"school/graph"
	"school/internal/auth"
	"school/internal/database"
	"school/internal/database/migrations"
	"school/internal/service"
	"school/proto/pb"
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

	migrations.Init()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))

	r := mux.NewRouter()

	r.Handle("/query", tokenMiddleware(srv))
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/metrics", promhttp.Handler())
	msHandler := handlers.LoggingHandler(os.Stdout, r)

	go grpcSchoolServer()
	go grpcSagaServer()

	log.Printf("SagaService is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, msHandler))
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

func grpcSagaServer() {
	lis, err := net.Listen("tcp", ":9095")
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

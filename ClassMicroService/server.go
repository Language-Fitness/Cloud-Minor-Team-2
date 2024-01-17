package main

import (
	"Class/graph"
	"Class/internal/auth"
	"Class/internal/service"
	"Class/proto/pb"
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
)

const defaultPort = "8082"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	graphQLServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))
	tokenMiddleware := auth.Middleware

	r := mux.NewRouter()
	r.Handle("/query", tokenMiddleware(graphQLServer))
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/metrics", promhttp.Handler())
	msHandler := handlers.LoggingHandler(os.Stdout, r)

	r.HandleFunc("/health/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	//migrations.Init()
	go grpcSagaServer()

	log.Printf("SagaService is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, msHandler))
}

func grpcSagaServer() {
	lis, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGRPCSagaServiceServer(grpcServer, service.NewSagaService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

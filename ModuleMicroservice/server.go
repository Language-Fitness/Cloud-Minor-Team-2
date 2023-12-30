package main

import (
	"Module/graph"
	"Module/internal/auth"
	"Module/internal/database"
	"Module/internal/service"
	"Module/proto/pb"
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

const defaultPort = "8081"

func main() {
	err := os.Setenv("GODEBUG", "http2debug=1")
	if err != nil {
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//migrations.Init()

	// Initialize GraphQL server
	graphQLServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))

	// Token middleware
	tokenMiddleware := auth.Middleware

	// Define your HTTP router
	r := mux.NewRouter()

	// GraphQL endpoint
	r.Handle("/query", tokenMiddleware(graphQLServer))

	// GraphQL playground
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// Metrics endpoint
	r.Handle("/metrics", promhttp.Handler())

	// Set up CORS middleware if needed
	// corsMiddleware := handlers.CORS(
	//     handlers.AllowedOrigins([]string{"*"}),
	//     handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	//     handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	// )

	// Create a handler with logging and CORS middleware
	msHandler := handlers.LoggingHandler(os.Stdout, r)
	// handler = corsMiddleware(handler)

	go grpcSagaServer()

	//// Start the HTTP server
	//log.Printf("SagaService is running on http://localhost:%s", port)
	//log.Fatal(http.ListenAndServe(":"+port, msHandler))

	log.Printf("SagaService is running on http://localhost:%s", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, msHandler))
}

func grpcSagaServer() {
	//// Initialize gRPC connection
	// TODO Dial is something that a grpc client uses to connect to a grpc server, but this seems to be a server
	//conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("failed to dial gRPC server: %v", err)
	//}
	//defer conn.Close()

	lis, err := net.Listen("tcp", ":9091") // Define your gRPC server address and port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC client using the connection
	// Create a new gRPC server instance

	grpcServer := grpc.NewServer()
	collection, _ := database.GetCollection()
	pb.RegisterGRPCSagaServiceServer(grpcServer, service.NewSagaService(collection))

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

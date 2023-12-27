package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"saga/graph"
	"saga/internal/auth"
	"saga/proto/pb"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8083"

func main() {
	err := os.Setenv("GODEBUG", "http2debug=2")
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

	//conn, err := grpc.Dial(os.Getenv("GRPC_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("failed to dial: %v", err)
	//}
	//defer conn.Close()
	//
	//// Create a gRPC client using the connection
	//grpcClient := pb.NewGRPCSagaServiceClient(conn)
	////migrations.Init()

	conn, err := grpc.Dial("host.docker.internal:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client using the connection
	client := pb.NewGRPCSagaServiceClient(conn)

	// Now you can use 'client' to make RPC calls to the gRPC server
	// For example:
	request := &pb.ObjectRequest{
		ObjectId:     "some_id",
		ObjectType:   pb.SagaObjectType_SCHOOL,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	response, err := client.FindObject(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to call FindObject RPC: %v", err)
	}

	fmt.Println(response)

	tokenMiddleware := auth.Middleware

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(client)}))

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/query", tokenMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

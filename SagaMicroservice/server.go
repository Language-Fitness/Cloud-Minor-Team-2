package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"saga/graph"
	"saga/internal/auth"
	"saga/proto/pb"
	"time"
)

const defaultPort = "8083"

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:9090", nil))
	}()
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

	// Unset HTTP proxy environment variable
	err = os.Unsetenv("HTTP_PROXY")
	if err != nil {
		fmt.Println("Error unsetting HTTP_PROXY:", err)
		return
	}

	// Unset HTTPS proxy environment variable
	err = os.Unsetenv("HTTPS_PROXY")
	if err != nil {
		fmt.Println("Error unsetting HTTPS_PROXY:", err)
		return
	}

	printEnvironment()

	//conn, err := grpc.Dial(os.Getenv("GRPC_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("failed to dial: %v", err)
	//}
	//defer conn.Close()
	//
	//// Create a gRPC client using the connection
	//grpcClient := pb.NewGRPCSagaServiceClient(conn)
	////migrations.Init()

	// If the user doesn't have a context with a deadline, create one
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fmt.Println("Creating context...")
	defer cancel()

	fmt.Println("Setting opts for gRPC dial...")
	opts := []grpc.DialOption{
		grpc.WithReturnConnectionError(), // Add the WithReturnConnectionError option
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	fmt.Println("Dialing gRPC server...")
	conn, err := grpc.DialContext(ctx, "host.docker.internal:9091", opts...)
	if err != nil {
		fmt.Printf("failed to dial gRPC server: %v\n", err)
		log.Printf("failed to dial gRPC server: %v", err)
	}
	fmt.Println("Dialing gRPC server...done")
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("error closing connection: %v\n", err)
			log.Printf("error closing connection: %v", err)
		}
	}(conn)

	conn, err = nil, fmt.Errorf("%v: %v", ctx.Err(), err)

	fmt.Println("Creating gRPC client...")
	// Create a gRPC client using the connection
	client := pb.NewGRPCSagaServiceClient(conn)

	// Now you can use 'client' to make RPC calls to the gRPC server
	// For example:
	request := pb.ObjectRequest{
		ObjectId:     "some_id",
		ObjectType:   pb.SagaObjectType_SCHOOL,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	fmt.Println("Calling FindObject RPC...")
	response, err := client.FindObject(context.Background(), &request)
	if err != nil {
		//log.Fatalf("failed to call FindObject RPC: %v", err)

		log.Printf("failed to call FindObject RPC: %v", err)
	}

	fmt.Println("\nResponse:", response)

	tokenMiddleware := auth.Middleware

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(client)}))

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/query", tokenMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func printEnvironment() {
	fmt.Println("Environment variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

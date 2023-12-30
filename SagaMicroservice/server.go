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

	fmt.Println("Setting opts for gRPC dial...")
	opts := []grpc.DialOption{
		grpc.WithReturnConnectionError(), // Add the WithReturnConnectionError option
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	fmt.Println("Dialing gRPC server...")
	conn, err := grpc.DialContext(context.Background(), "host.docker.internal:9091", opts...)
	if err != nil {
		fmt.Printf("failed to dial gRPC server: %v\n", err)
		log.Printf("failed to dial gRPC server: %v", err)
	}
	fmt.Println("Dialing gRPC server...done")
	defer conn.Close()

	fmt.Println("Creating gRPC client...")
	// Create a gRPC client using the connection
	client := pb.NewGRPCSagaServiceClient(conn)

	// Now you can use 'client' to make RPC calls to the gRPC server
	// For example:
	request := pb.ObjectRequest{
		BearerToken:  "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJIaUpNcWZhTGFWQXBiME5JTEpweTlacmdtRzBERElIaWpVZklVWjM2NXJvIn0.eyJleHAiOjE3MDM5NDc5NzMsImlhdCI6MTcwMzk0NzY3MywianRpIjoiNjk1ZWNlNTktZDYwOS00MjExLWFiMjQtNjA2ODMwMTFkMTQ0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjZiMDNiYTVkLTVkMGUtNGRkOC05ZjdmLTkyOGU3NWVhOGVjYSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIzYjYxNGNmNy00NmVjLTQ5NDEtOWU3Zi0wODkzZGRiODA3NmUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZ2V0X2NsYXNzZXNfYWxsIiwidXBkYXRlX3NjaG9vbCIsImZpbHRlcl9zY2hvb2xfbWFkZV9ieSIsImZpbHRlcl9zY2hvb2xfbmFtZSIsImZpbHRlcl9tb2R1bGVfY2F0ZWdvcnkiLCJmaWx0ZXJfY2xhc3NfbWFkZV9ieSIsImZpbHRlcl9tb2R1bGVfc29mdERlbGV0ZSIsImdldF9leGVyY2lzZXMiLCJnZXRfY2xhc3NlcyIsImRlbGV0ZV9tb2R1bGUiLCJkZWxldGVfZXhlcmNpc2UiLCJnZXRfc2Nob29scyIsInVwZGF0ZV9leGVyY2lzZSIsImdldF9leGVyY2lzZSIsImRlbGV0ZV9tb2R1bGVfYWxsIiwiY3JlYXRlX2V4ZXJjaXNlIiwiZ2V0X3NjaG9vbCIsImRlbGV0ZV9leGVyY2lzZV9hbGwiLCJmaWx0ZXJfc2Nob29sX2xvY2F0aW9uIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJmaWx0ZXJfbW9kdWxlX2RpZmZpY3VsdHkiLCJjcmVhdGVfbW9kdWxlIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwidXBkYXRlX2V4ZXJjaXNlX2FsbCIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX3NvZnREZWxldGUiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImdldF9tb2R1bGVzX2FsbCIsImZpbHRlcl9jbGFzc19tb2R1bGVfaWQiLCJmaWx0ZXJfbW9kdWxlX3NjaG9vbF9pZCIsImZpbHRlcl9tb2R1bGVfbWFkZV9ieSIsImZpbHRlcl9jbGFzc19uYW1lIiwidXBkYXRlX2NsYXNzX2FsbCIsImZpbHRlcl9tb2R1bGVfbmFtZSIsInVwZGF0ZV9tb2R1bGUiLCJnZXRfY2xhc3MiLCJkZWxldGVfc2Nob29sX2FsbCIsImZpbHRlcl9tb2R1bGVfcHJpdmF0ZSIsInVwZGF0ZV9jbGFzcyIsImdldF9zY2hvb2xzX2FsbCIsImZpbHRlcl9jbGFzc19zb2Z0RGVsZXRlIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiM2I2MTRjZjctNDZlYy00OTQxLTllN2YtMDg5M2RkYjgwNzZlIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.ufqx_RD2A5abIdqEvf79lwl9bsQjZnIga54v82OdEKDpH47IF4yfnKIwl5f4sCpVeyaPl_ihPpYLtBwuD2ZLD-O-u6zRnFyVm3sXuAeN2CC3FOEWZtxr0gxECySaW7k3Oj7AWZimn_yxJfxyElRuNhlg4811gFJ1bZgGkl_3vJvg_61FEIBQB74vQA51jx27Y2-kSxdSMxXAkgWVNYjFtjgDyzeGZUHibqw8uLX4NYASprW4lGDVu-A3S_Vj3dJvJJON6Oe_8-IS-LH2Vw6olJNjEonxm9x5HJAWwUcn_Md4ShUB3u-k9jT1MAFkke1p4h5wuRnTa5mY3yjaF8LiRw",
		ObjectId:     "d3a22799-9e85-4dfb-ae53-a760b0f6314b",
		ObjectType:   pb.SagaObjectType_SCHOOL,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	fmt.Println("Calling FindObject RPC...")
	response, err := client.FindObject(context.Background(), &request)
	if err != nil {
		//log.Fatalf("failed to call FindObject RPC: %v", err)
		log.Printf("failed to call FindObject RPC: %v", err)
	}

	fmt.Println("\nResponse 1:", response.Objects[0].ObjectId)
	fmt.Println("\nResponse 2:", response.Objects[0].ObjectStatus)
	fmt.Println("\nResponse 3:", response.Objects[0].ObjectType)

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

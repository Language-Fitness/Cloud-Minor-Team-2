package helper

import (
	"ExerciseMicroservice/graph"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const defaultPort = "8084"

func CreateClient() *client.Client {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	// Create the gqlgen client for testing
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))
	c := client.New(h)

	return c
}

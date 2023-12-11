package main

import (
	"Module/graph"
	"Module/internal/auth"
	"Module/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	randomModules := GenerateRandomModules(1000)
	fmt.Println(randomModules)
	collection, _ := database.GetCollection()
	// Insert each random module into the collection
	for _, module := range randomModules {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout

		_, err := collection.InsertOne(ctx, module)
		cancel() // Cancel the context after the operation is done
		if err != nil {
			log.Fatal(err)
		}
	}

	tokenMiddleware := auth.Middleware

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/query", tokenMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type Category string

const (
	CategoryGrammatica            Category = "Grammatica"
	CategorySpelling              Category = "Spelling"
	CategoryWoordenschat          Category = "Woordenschat"
	CategoryUitdrukkingen         Category = "Uitdrukkingen"
	CategoryInterpunctie          Category = "Interpunctie"
	CategoryWerkwoordvervoegingen Category = "Werkwoordvervoegingen"
	CategoryFastTrack             Category = "Fast_Track"
)

type LanguageLevel string

const (
	LanguageLevelA1 LanguageLevel = "A1"
	LanguageLevelA2 LanguageLevel = "A2"
	LanguageLevelB1 LanguageLevel = "B1"
	LanguageLevelB2 LanguageLevel = "B2"
	LanguageLevelC1 LanguageLevel = "C1"
	LanguageLevelC2 LanguageLevel = "C2"
)

type Module struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  LanguageLevel `json:"difficulty"`
	Category    Category      `json:"category"`
	MadeBy      string        `json:"made_by"`
	Private     bool          `json:"private"`
	Key         *string       `json:"key,omitempty"`
	CreatedAt   *string       `json:"created_at,omitempty"`
	UpdatedAt   *string       `json:"updated_at,omitempty"`
	SoftDeleted *bool         `json:"soft_deleted,omitempty"`
}

func GenerateRandomModules(n int) []Module {
	var modules []Module
	for i := 0; i < n; i++ {
		module := Module{
			ID:          uuid.New().String(),
			Name:        fmt.Sprintf("Module%d", i),
			Description: fmt.Sprintf("Description for Module%d", i),
			Difficulty:  LanguageLevelA1, // Random difficulty between 1 and 5
			Category:    CategoryGrammatica,
			MadeBy:      uuid.New().String(),  // Random UUID for MadeBy
			Private:     false,                // Randomly set to true or false
			Key:         generateRandomHash(), // Random hash for Key
			CreatedAt:   generateRandomDate(), // Random date for CreatedAt
			UpdatedAt:   generateRandomDate(), // Random date for UpdatedAt
			SoftDeleted: new(bool),            // Set to false
		}
		*module.SoftDeleted = false
		modules = append(modules, module)
	}
	return modules
}

// generateRandomHash generates a random hash of 15 letters
func generateRandomHash() *string {
	hash := make([]byte, 15)
	rand.Seed(time.Now().UnixNano())
	for i := range hash {
		hash[i] = byte(rand.Intn(26) + 65) // ASCII A-Z
	}
	result := string(hash)
	return &result
}

// generateRandomDate generates a random date string
func generateRandomDate() *string {
	rand.Seed(time.Now().UnixNano())
	days := rand.Intn(365) // Random number of days
	randomDate := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)
	return &randomDate
}

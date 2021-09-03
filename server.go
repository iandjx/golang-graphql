package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/iandjx/hackernews/graph"
	"github.com/iandjx/hackernews/graph/generated"
	"github.com/iandjx/hackernews/internal/auth"
	database "github.com/iandjx/hackernews/internal/pkg/db/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	// database.InitDB()
	// database.Migrate()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

package main

import (
	"fmt"
	"github.com/cass-dlcm/pomodoro_tasks/backend/auth"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cass-dlcm/pomodoro_tasks/graph"
	"github.com/cass-dlcm/pomodoro_tasks/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.InitDB()

	auth.InitAuth()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	buildHandler := http.FileServer(http.Dir("frontend/build"))

	r := http.NewServeMux()
	r.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", auth.JWTMiddleware(srv))
	r.Handle("/", buildHandler)

	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
	}

	log.Fatal(server.ListenAndServe())
}

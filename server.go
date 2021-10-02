package main

import (
	"fmt"
	"github.com/cass-dlcm/pomodoro_tasks/backend/auth"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cass-dlcm/pomodoro_tasks/graph"
	"github.com/cass-dlcm/pomodoro_tasks/graph/generated"
)

const defaultPort = "8080"
const FSPATH = "./frontend/build/"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := db.InitDB(); err != nil {
		log.Panicln(err)
	}

	if err := auth.InitAuth(); err != nil {
		log.Panicln(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	fs := http.FileServer(http.Dir(FSPATH))

	r := http.NewServeMux()
	r.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", auth.JWTMiddleware(srv))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If the requested file exists then return if; otherwise return index.html (fileserver default page)
		if r.URL.Path != "/" {
			fullPath := FSPATH + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
			_, err := os.Stat(fullPath)
			if err != nil {
				if !os.IsNotExist(err) {
					panic(err)
				}
				// Requested file does not exist so we return the default (resolves to index.html)
				r.URL.Path = "/"
			}
		}
		fs.ServeHTTP(w, r)
	})

	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
	}

	log.Fatalln(server.ListenAndServe())
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/marukhno/go-graphql/graph"
	"github.com/marukhno/go-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
	log.Printf("connect to https://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServeTLS(":"+port, "certs/go-graphql.crt", "certs/go-graphql.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/HydroOSS/HydroAPI/graph"
	"github.com/HydroOSS/HydroAPI/graph/generated"
)

func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	rethinkAddr := "localhost:28015"
	if envAddr := os.Getenv("DB_ADDRESS"); envAddr != "" {
		rethinkAddr = envAddr
	}

	session, err := r.Connect(r.ConnectOpts{
		Address:    rethinkAddr,
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		log.Fatalln(err)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: session,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("Attempting to start server at http://localhost:%s/\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

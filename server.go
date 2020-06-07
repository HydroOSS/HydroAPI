package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
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

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Infof("Attempting to start server at http://localhost:%s/\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"ElTube_BackEnd/graph"
	"ElTube_BackEnd/graph/generated"
	"ElTube_BackEnd/postgre"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "2400"

func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	db := postgre.New(&pg.Options{
		Addr:     ":5432",
		User:     "pdmlvtwtxvnkkk",
		Password: "74ddcb4d854afe215675d3b3bcde99304b147507487b8b19f2e07e1c2ea8c931",
		Database: "dco107usf3flks",
	})

	option, err := pg.ParseURL("postgres://pdmlvtwtxvnkkk:74ddcb4d854afe215675d3b3bcde99304b147507487b8b19f2e07e1c2ea8c931@ec2-18-232-143-90.compute-1.amazonaws.com:5432/dco107usf3flks")
	if err != nil {
		panic(err);
	}

	db = pg.Connect(option);

	db.AddQueryHook(postgre.DBLogger{})

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

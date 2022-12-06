package main

import (
	"log"
	"net/http"

	"database/sql"

	api "github.com/anmi/go-ttsa/api"
	q "github.com/anmi/go-ttsa/db"
	apiService "github.com/anmi/go-ttsa/service"
	utils "github.com/anmi/go-ttsa/utils"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	// Create service instance.
	db, err := sql.Open("postgres", "postgres://myusername:mypassword@127.0.0.1:5432/tinytracker?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	queries := q.New(db)
	service := &apiService.TTApiService{Queries: queries}
	// Create generated server.
	srv, err := api.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})
	withRAR := utils.RARPropagator{Next: srv}
	withCors := c.Handler(withRAR)
	if err := http.ListenAndServe(":8080", withCors); err != nil {
		log.Fatal(err)
	}
}

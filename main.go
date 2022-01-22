package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"green/database"
	"green/handlers/mazeHandler"
	"log"
	"net/http"
	"os"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func routers() *chi.Mux {
	router.Post("/maze", mazeHandler.InsertMaze())
	router.Delete("/maze/{id}", mazeHandler.DeleteMaze())
	router.Put("/maze/{id}", mazeHandler.UpdateMaze())
	router.Get("/maze/{id}", mazeHandler.GetMaze())
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("It's up!"))
	})
	return router

}

func main() {
	routers()
	database.InitDB()

	port := os.Getenv("PORT")
	defaultPort := "10000"

	if !(port == "") {
		log.Fatal(http.ListenAndServe(":"+port, router))
	} else {
		log.Printf("Starting up on http://localhost:%s", defaultPort)
		log.Fatal(http.ListenAndServe(":"+defaultPort, router))
	}
}

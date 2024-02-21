package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("SERVER_PORT")

	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	chiRouter.Get("/", home)

	chiRouter.Route("/snippet", func(cr chi.Router) {

		cr.Get("/view", snippetView)

		cr.Post("/create", snippetCreate)
	})

	chiRouter.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Route does not exist...\n"))
	})

	chiRouter.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		snippetCreateMethodNotAllowed(w, r)

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method is not valid\n"))
	})

	log.Println("Starting server on :" + PORT)
	http.ListenAndServe(":"+PORT, chiRouter)
}

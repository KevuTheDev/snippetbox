package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load Environment Variables
	errGDE := godotenv.Load()
	if errGDE != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("SERVER_PORT")

	// Setup Loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	// Setup Router
	mux := http.NewServeMux()

	// - Static Files
	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))

	// - Routing
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("Starting server on port %s", PORT)
	errServer := http.ListenAndServe(":"+PORT, mux)
	errorLog.Fatal(errServer)
}

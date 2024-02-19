package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func tesst(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		w.Write([]byte("Create a new snippet...\n"))
		return errors.New(q)
	}

	log.Println("ERROR POST")
	w.Write([]byte("Create a new snippet...\n"))
	return nil
}

// func (t *Template) Render(w io.Writer, title string, data interface{}, c echo.Context) error {
// 	return t.tmpl.ExecuteTemplate(w, title, data)
// }

func main() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("SERVER_PORT")

	chiR := chi.NewRouter()
	chiR.Use(middleware.Logger)

	chiR.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Snippetbox\n"))
	})

	chiR.Route("/snippet", func(rr chi.Router) {
		rr.Get("/view", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Display a specific snippet...\n"))

		})

	})

	log.Println("Starting server on :" + PORT)
	http.ListenAndServe(":"+PORT, chiR)
}

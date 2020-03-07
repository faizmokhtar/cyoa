package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/faizmokhtar/cyoa/cyoa"
	"github.com/gorilla/mux"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.Handle("/{key}", NewHandler(story))
	r.Handle("/", NewHandler(story))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

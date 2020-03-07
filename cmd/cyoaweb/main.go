package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/faizmokhtar/cyoa"
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

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3030",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

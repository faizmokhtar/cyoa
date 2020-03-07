package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/faizmokhtar/cyoa/cyoa"
	"github.com/gorilla/mux"
)

type handlers struct {
	s cyoa.Story
}

func (h handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/page.gohtml"))
	path := mux.Vars(r)["key"]
	if path == "" || path == "/" {
		path = "intro"
	}

	if chapter, ok := h.s[path]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found", http.StatusNotFound)
}

func NewHandler(s cyoa.Story) http.Handler {
	return handlers{s}
}

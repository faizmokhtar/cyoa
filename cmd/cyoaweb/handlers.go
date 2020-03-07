package main

import (
	"cyoa"
	"net/http"
	"text/template"
)

type handlers struct {
	s cyoa.Story
}

func (h handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/page.gohtml"))
	err := tmpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

func NewHandler(s cyoa.Story) http.Handler {
	return handlers{s}
}

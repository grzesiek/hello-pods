package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	v1 "k8s.io/api/core/v1"
)

type Server struct {
	handler Handler
}

func (s *Server) Serve(f func(h Handler) http.HandlerFunc) {
	if err := http.ListenAndServe(":8080", f(s.handler)); err != nil {
		log.Fatal(err)
	}
}

type Handler struct {
	index string
}

func (h *Handler) ServeError(err error) http.HandlerFunc {
	log.Print(err)

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Error occured: %q", err)
	}
}

func (h *Handler) ServePods(pods []v1.Pod) http.HandlerFunc {
	html, err := ioutil.ReadFile(h.index)
	if err != nil {
		return h.ServeError(err)
	}

	tpl, err := template.New("pods").Parse(string(html))
	if err != nil {
		return h.ServeError(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, pods)
	}
}

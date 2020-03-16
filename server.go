package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	index string
}

func (s *Server) ServeError(err error) {
	log.Print(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Error occured: %q", err)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) ServePods(pods *Pods) {
	html, err := ioutil.ReadFile(s.index)
	if err != nil {
		s.ServeError(err)
		return
	}

	tpl, err := template.New("todos").Parse(string(html))
	if err != nil {
		s.ServeError(err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, pods)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

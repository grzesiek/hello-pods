package main

import (
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	server := &Server{handler: Handler{index: "index.html.tpl"}}

	server.Serve(func(handler Handler) http.HandlerFunc {
		config, err := rest.InClusterConfig()
		if err != nil {
			return handler.ServeError(err)
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			return handler.ServeError(err)
		}

		pods, err := Pods{client: clientset, namespace: "default"}.List()
		if err != nil {
			return handler.ServeError(err)
		}

		return handler.ServePods(pods)
	})
}

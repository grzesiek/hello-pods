package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	server := &Server{index: "index.html.tpl"}

	config, err := rest.InClusterConfig()
	if err != nil {
		server.ServeError(err)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		server.ServeError(err)
		return
	}

	server.ServePods(&Pods{client: clientset, namespace: "default"})
}

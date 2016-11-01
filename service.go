package main

import (
	"github.com/karlmoad/k8s-example/service"
	"github.com/karlmoad/k8s-example/service/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	var routes = service.Routes{
		service.Route {
			"Echo",
			"POST",
			"/echo",
			handlers.Echo,
		},
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := service.NewRouter(routes)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

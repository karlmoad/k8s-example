package main

import (
	"github.com/karlmoad/k8s-example/service"
	"github.com/karlmoad/k8s-example/service/handlers"
	"log"
	"net/http"
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

	router := service.NewRouter(routes)

	log.Fatal(http.ListenAndServe(":8080", router))
}

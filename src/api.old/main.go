// Package main entrypoint to the API
//
// Documentation for Excalibur API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"excalibur/internal/handler"
	"excalibur/internal/repository"
	"excalibur/internal/routerbuilder"
	"excalibur/internal/services"
	"log"
	"net/http"
	"time"
)

func main() {
	l := log.Default()
	dao := repository.NewDAO(*l)

	serviceCollection := services.NewServiceCollection(dao, *l)
	handlerCollection := handler.NewHandlerCollection(serviceCollection, *l)

	srv := createServer(handlerCollection)

	log.Fatal(srv.ListenAndServe())
}

func createServer(hc handler.Collection) *http.Server {
	rb := routerbuilder.New(hc)
	r := rb.Init()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}

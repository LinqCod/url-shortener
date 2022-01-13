package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	router := initRoutes()

	server := http.Server{
		Addr:           ":8080",
		Handler:        router,
		WriteTimeout:   5 * time.Second,
		ReadTimeout:    5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		panic("Error while starting server")
	}
	log.Fatalf("Server is running on port 8080")
}

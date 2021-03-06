package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/johnpmitsch/verbugata-go-api/config"
	"github.com/johnpmitsch/verbugata-go-api/router"
)

func main() {
	appConf := config.AppConfig()
	appRouter := router.New()
	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

// Greets user
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

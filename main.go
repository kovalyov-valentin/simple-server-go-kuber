package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kovalyov-valentin/simple-server-go-kuber/handlers"
	"github.com/kovalyov-valentin/simple-server-go-kuber/version"
)

func main() {
	log.Printf(
		"Srart the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set")
	}

	r := handlers.Router(version.BuildTime, version.Commit, version.Release)
	interput := make(chan os.Signal, 1)
	signal.Notify(interput, os.Interrupt, syscall.SIGTERM)
	srv := &http.Server {
		Addr : ":" + port, 
		Handler: r,
	}

	go func () {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Print("The service is ready to listen and serve.")
		killSignal := <- interput
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}

	
	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Printf("Down")
}

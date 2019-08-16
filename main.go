package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	config "ursho/config"
	handler "ursho/handler"
	mongodb "ursho/storage/mongodb"
)

func main() {
	configPath := flag.String("config", "./config/config.json", "path of the config file")

	flag.Parse()
	// Read config
	config, err := config.FromFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	svc, err := mongodb.New(config.Mongo.Host, config.Mongo.Port, config.Mongo.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer svc.Close()

	// Create a server
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		Handler: handler.New(config.Options.Prefix, svc),
		}

	// Check for a closing signal
	go func() {
		// Graceful shutdown
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		sig := <-sigquit
		log.Printf("caught sig: %+v", sig)
		log.Printf("Gracefully shutting down server...")

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Unable to shut down server: %v", err)
		} else {
			log.Println("Server stopped")
		}
	}()

	// Start server
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}

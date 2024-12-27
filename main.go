package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"beto0607.com/blober/src/config"
	"beto0607.com/blober/src/data"
	"beto0607.com/blober/src/router"
)

func main() {
	data.ConnectToDB()

	server := initServer()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("Starting server on " + server.Addr)
		server.ListenAndServe()
	}()

	<-ctx.Done()

	server.Shutdown(context.TODO())
	data.DisconnectDB()
	log.Println("Server shutdown")
	log.Println("final")
}

func initServer() *http.Server {
	serverPort, err := config.GetEnvVar("PORT")
	if err != nil {
		serverPort = "8978"
	}
	hostname, err := config.GetEnvVar("HOST")
	if err != nil {
		hostname = "api.blober.local"
	}

	router := router.Route()

	serverAddress := hostname + ":" + serverPort

	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  config.DefaultReadTimeout,
		WriteTimeout: config.DefaultWriteTimeout,
		Handler:      router,
	}

	return server
}

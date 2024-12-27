package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"beto0607.com/blober/src/core"
)

func main() {
	core.ConnectToDB()

	server := core.InitServer()

	err := core.InitFS()

	if err != nil {
		log.Fatalln("Couldn't init FS")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("Starting server on " + server.Addr)
		server.ListenAndServe()
	}()

	<-ctx.Done()

	server.Shutdown(context.TODO())
	core.DisconnectDB()
	log.Println("Server shutdown")
	log.Println("final")
}

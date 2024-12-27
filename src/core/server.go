package core

import (
	"net/http"

	"beto0607.com/blober/src/config"
)

func InitServer() *http.Server {
	serverPort, err := config.GetEnvVar("PORT")
	if err != nil {
		serverPort = "8978"
	}
	hostname, err := config.GetEnvVar("HOST")
	if err != nil {
		hostname = "api.blober.local"
	}

	router := InitRouting()

	serverAddress := hostname + ":" + serverPort

	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  config.DefaultReadTimeout,
		WriteTimeout: config.DefaultWriteTimeout,
		Handler:      router,
	}

	return server
}

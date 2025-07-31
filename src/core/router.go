package core

import (
	"log"
	"net/http"

	blob_slice "beto0607.com/blober/src/slices/blob"
	container_slice "beto0607.com/blober/src/slices/container"
	health_slice "beto0607.com/blober/src/slices/health"
)

func InitRouting() *http.ServeMux {
	apiRouter := doApiRouting()
	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))
	return router
}

func doApiRouting() *http.ServeMux {
	log.Println("Doing API routing...")
	router := http.NewServeMux()
	health_slice.HealthAPIRouting(router)
	blob_slice.BlobAPIRouting(router)
	container_slice.ContainerAPIRouting(router)

	return router
}

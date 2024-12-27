package router

import (
	"log"
	"net/http"

	blob_slice "beto0607.com/blober/src/slices/blob"
	container_slice "beto0607.com/blober/src/slices/container"
)

func Route() *http.ServeMux {
	apiRouter := doApiRouting()
	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))
	return router
}

func doApiRouting() *http.ServeMux {
	log.Println("Doing API routing...")
	router := http.NewServeMux()
	blob_slice.BlobAPIRouting(router)
	container_slice.ContainerAPIRouting(router)

	return router
}

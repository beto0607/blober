package router

import (
	"log"
	"net/http"

	"beto0607.com/blober/src/controllers"
	blob_slice "beto0607.com/blober/src/slices/blob"
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
	containerAPIRouting(router)

	return router
}

func containerAPIRouting(router *http.ServeMux) {
	router.HandleFunc("GET /containers/{id}/metadata", controllers.GetContainerMetadata)
	router.HandleFunc("GET /containers/{id}/list", controllers.ListContainerBlob)
	router.HandleFunc("POST /containers", controllers.PostContainer)
	router.HandleFunc("PUT /containers/{id}", controllers.PutContainer)
	router.HandleFunc("DELETE /containers/{id}", controllers.DeleteContainer)
	log.Println("Container API added")
}

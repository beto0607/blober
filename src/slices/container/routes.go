package container_slice

import (
	"log"
	"net/http"
)

func ContainerAPIRouting(router *http.ServeMux) {
	router.HandleFunc("GET /containers/{id}/metadata", GetContainerMetadata)
	router.HandleFunc("GET /containers/{id}/list", ListContainerBlob)
	router.HandleFunc("POST /containers", PostContainer)
	router.HandleFunc("PUT /containers/{id}", PutContainer)
	router.HandleFunc("DELETE /containers/{id}", DeleteContainer)
	log.Println("Container API added")
}

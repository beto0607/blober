package blob_slice

import (
	"log"
	"net/http"
)

func BlobAPIRouting(router *http.ServeMux) {
	router.HandleFunc("GET /blobs/{id}", GetBlob)
	router.HandleFunc("GET /blobs/{id}/metadata", GetBlobMetadata)
	router.HandleFunc("POST /blobs", PostBlob)
	router.HandleFunc("PUT /blobs/{id}", PutBlob)
	router.HandleFunc("DELETE /blobs/{id}", DeleteBlob)
	log.Println("Blob API added")
}

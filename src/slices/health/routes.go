package health

import (
	"log"
	"net/http"
)

func HealthAPIRouting(router *http.ServeMux) {
	router.HandleFunc("GET /health", GetHealth)
	log.Println("Health API added")
}

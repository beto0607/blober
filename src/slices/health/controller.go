package health

import (
	"encoding/json"
	"net/http"
	"time"
)

var start = time.Now()

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Since(start).String(),
	})
}

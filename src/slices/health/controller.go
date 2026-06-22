package health

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var since = time.Now().UTC()

func GetHealth(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC()
	entity := map[string]string{
		"status":      "ok",
		"since":       since.Format(time.RFC3339),
		"currentTime": now.Format(time.RFC3339),
		"aliveFor":    strconv.FormatInt((now.UnixMilli()-since.UnixMilli())/1000, 10),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&entity)
}

package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept-Encoding", "gzip, deflate, br")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
package responses

import (
	"encoding/json"
	"net/http"
)

func OK(w http.ResponseWriter, dto any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	bytes, _ := json.Marshal(dto)

	w.Write(bytes)
}

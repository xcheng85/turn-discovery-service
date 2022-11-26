package utils

import (
	"encoding/json"
	"net/http"
)

func WriteStatusResponse(w http.ResponseWriter, statusCode int, message string) {
	type statusResponse struct {
		Status string
	}
	status := statusResponse{
		Status: message,
	}
	WriteJSONResponse(w, statusCode, status)
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	enc := json.NewEncoder((w))
	enc.SetIndent("", " ")
	return enc.Encode(data)
}

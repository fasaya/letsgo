package helpers

import (
	"encoding/json"
	"golang-for-women/sesi-6/types"
	"net/http"
)

func HandleErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	response := types.Response{
		Error:   true,
		Message: message,
		Data:    nil,
	}
	json.NewEncoder(w).Encode(response)
}

func HandleSuccessfulResponse(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	response := types.Response{
		Error:   false,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}

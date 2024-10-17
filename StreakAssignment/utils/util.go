package utils

import (
	"encoding/json"
	"net/http"
	//"net/http"
)

//SendRESPONSE sends a JSON response to the client......
func SendResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// EncodING data
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
}
}

// SendError sends an error response to the client....
func SendError(w http.ResponseWriter, message string, status int) {
	SendResponse(w, map[string]string{"error": message}, status)
}

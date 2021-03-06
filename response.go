package gohttputil

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// RespondJSON sends a JSON object as an HTTP response.
// If the object cannot be marshaled, an error message will be sent and
// false will be returned.
// Otherwise, true is returned.
func RespondJSON(res http.ResponseWriter, code int, obj interface{}) bool {
	res.Header().Set("Content-Type", "application/json")
	if marshaled, err := json.Marshal(obj); err == nil {
		res.Header().Set("Content-Length", strconv.Itoa(len(marshaled)))
		res.WriteHeader(code)
		res.Write(marshaled)
		return true
	} else {
		data := []byte("\"Failed to encode object\"")
		res.Header().Set("Content-Length", strconv.Itoa(len(data)))
		// Preserve error code (if there is one)
		if code == 200 {
			res.WriteHeader(http.StatusInternalServerError)
		} else {
			res.WriteHeader(code)
		}
		res.Write(data)
		return false
	}
}
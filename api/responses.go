package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Marshals and sends the given payload as JSON or status code 500 if that fails.
// Any occurring errors are logged.
func SendJson(w http.ResponseWriter, payload interface{}) {
	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Failed to marshal data")
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(b) // writes status 200 implicitly
	if err != nil {
		fmt.Println("Failed to write response")
	}
}

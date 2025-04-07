package errs

import (
	"encoding/json"
	"net/http"
	"time"
)

func Throw(w http.ResponseWriter, err HttpErrorInterface, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode())

	json.NewEncoder(w).Encode(map[string]interface{}{
		"timestamp":  time.Now().Format(time.RFC3339),
		"message":    err.Error(),
		"type":       err.Type(),
		"statusCode": err.StatusCode(),
		"path":       r.URL.Path,
	})
}

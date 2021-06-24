package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string
}

func WriteError(w http.ResponseWriter, statusCode int, error error) {
	w.WriteHeader(statusCode)
	if error != nil {
		err := Error{Message: error.Error()}
		json.NewEncoder(w).Encode(err)
	}
}
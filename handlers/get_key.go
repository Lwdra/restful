package handlers

import (
	"net/http"

	"fmt"

	"github.com/restful/storage"
)

// GetKey returns a httpHandler that can Get a key from the db
func GetKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {

		key := req.URL.Query().Get("key")
		if key == "" {
			http.Error(respWriter, "no key name provided", http.StatusBadRequest)
			return
		}

		val, err := db.Get(key)
		if err == storage.ErrNotFound {
			http.Error(respWriter, "Not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(respWriter, fmt.Sprintf("Unknown error: %s", err), http.StatusInternalServerError)
			return
		}

		respWriter.WriteHeader(http.StatusOK)
		respWriter.Write(val)
		return
	})
}

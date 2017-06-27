package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/restful/storage"
)

// PutKey returns a httpHandler that can set a value in the db
func PutKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {

		key := req.URL.Query().Get("key")
		if key == "" {
			http.Error(respWriter, "no key name provided", http.StatusBadRequest)
			return
		}

		defer req.Body.Close()

		val, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(respWriter, "Error reading body", http.StatusBadRequest)
			return
		}
		if err := db.Set(key, val); err != nil {
			http.Error(respWriter, "Error setting value", http.StatusInternalServerError)
			return
		}

		respWriter.WriteHeader(http.StatusOK)
		return
	})
}

package middleware

import (
	"assignment/utils"
	"encoding/json"
	"net/http"
)

func ErrorHandlerMiddleware(next func (w http.ResponseWriter, r *http.Request) *http_utils.HttpError) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			(*err).Path = r.URL.Path
			w.WriteHeader(http.StatusBadRequest)
			encoder := json.NewEncoder(w)
			encoder.Encode(*err)
		}
	})
}

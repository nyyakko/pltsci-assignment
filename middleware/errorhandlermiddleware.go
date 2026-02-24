package middleware

import (
	"encoding/json"
	"net/http"
)

func ErrorHandlerMiddleware(next func (w http.ResponseWriter, r *http.Request) error) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			var res struct { Error string `json:"error"` }
			res.Error = err.Error()
			encoder := json.NewEncoder(w)
			encoder.Encode(res)
		}
	})
}

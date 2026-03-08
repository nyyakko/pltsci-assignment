package main

import (
	"assignment/api/hoover"
	"assignment/middleware"
	"assignment/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func checkHealth(w http.ResponseWriter, r *http.Request) *http_utils.HttpError {
	encoder := json.NewEncoder(w)
	encoder.Encode(struct { Status string `json:"status"` }{ Status: "ok" })
	return nil
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /v1/health", middleware.ErrorHandlerMiddleware(checkHealth))
	mux.Handle("POST /v1/cleaning-sessions", middleware.ErrorHandlerMiddleware(hoover.Controller.CleaningSessions))

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", mux)
}

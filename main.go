package main

import (
	"assignment/api/hoover"
	"assignment/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("POST /v1/cleaning-sessions", middleware.ErrorHandlerMiddleware(hoover.Controller.CleaningSessions))

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", mux)
}

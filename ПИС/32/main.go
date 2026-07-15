package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func userHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		Response{
			Code:    200,
			Message: "Hello from User Handler",
		},
	)
}

func adminHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		Response{
			Code:    200,
			Message: "Hello from Admin Handler",
		},
	)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler).Queries("type", "user")
	r.HandleFunc("/users", adminHandler).Queries("type", "admin")

	log.Fatal(
		http.ListenAndServe(":3000", r),
	)
}

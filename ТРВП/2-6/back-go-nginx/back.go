package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Message: "Hello from back-go",
	})
}

func main() {
	http.HandleFunc("/api", handler)

	log.Println("Server started on :3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}

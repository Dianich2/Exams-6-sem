package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	var test Test
	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil || test.Id == 0 || test.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			ErrorResponse{
				Code:    400,
				Message: "Bad Request",
			},
		)
		return
	}

	json.NewEncoder(w).Encode(test)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tests", handler).Methods(http.MethodPost)

	log.Fatal(
		http.ListenAndServe(":3000", r),
	)
}

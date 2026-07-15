package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func handlerA(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(
		Response{
			Code:    200,
			Message: fmt.Sprintf("Method: %s, path: %s", r.Method, r.URL.Path),
		},
	)
}

func handlerB(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	var test Test
	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil || test.Id == 0 || test.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			Response{
				Code:    400,
				Message: "Bad Request",
			},
		)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(
		Response{
			Code:    201,
			Message: fmt.Sprintf("Id: %d, Name: %s", test.Id, test.Name),
		},
	)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/A", handlerA).Methods(http.MethodGet)
	r.HandleFunc("/B", handlerB).Methods(http.MethodPost)
	log.Fatal(
		http.ListenAndServe(":3000", r),
	)
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var test Test
		err := json.NewDecoder(r.Body).Decode(&test)
		if err != nil || test.Name == "" || test.Id == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				ErrorResponse{
					Code:    400,
					Message: "Bad Request",
				},
			)
			return
		}
		json.NewEncoder(w).Encode(&test)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(
			ErrorResponse{
				Code:    405,
				Message: "Method Not Allowed",
			},
		)
	}
}

func main() {
	http.HandleFunc("/tests", handler)
	log.Printf("Server start on 3000\n")
	log.Fatal(
		http.ListenAndServe(":3000", nil),
	)
}

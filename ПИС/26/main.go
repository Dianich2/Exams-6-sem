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

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			Response{
				Code:    400,
				Message: err.Error(),
			},
		)
		return
	}
	resp := map[string]string{
		"id":   r.FormValue("id"),
		"name": r.FormValue("name"),
	}
	if resp["id"] == "" || resp["name"] == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			Response{
				Code:    400,
				Message: "Bad Request",
			},
		)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		json.NewEncoder(w).Encode(
			Response{
				Code:    500,
				Message: err.Error(),
			},
		)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tests", handler).Methods(http.MethodPost)

	log.Fatal(
		http.ListenAndServe(":3000", r),
	)
}

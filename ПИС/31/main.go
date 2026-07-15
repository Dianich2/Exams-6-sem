package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			Response{
				Code:    400,
				Message: "Bad Request",
			},
		)
		return
	}

	x := r.FormValue("x")
	y := r.FormValue("y")
	file, header, err := r.FormFile("s")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			Response{
				Code:    400,
				Message: "Bad Request: cannot get file field",
			},
		)
		return
	}
	defer file.Close()
	dist, err := os.Create(header.Filename)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(
			Response{
				Code:    500,
				Message: "Server error",
			},
		)
		return
	}

	defer dist.Close()

	if _, err := io.Copy(dist, file); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(
			Response{
				Code:    500,
				Message: "Server error",
			},
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(
		Response{
			Code:    200,
			Message: fmt.Sprintf("x = %s, y = %s, filename = %s", x, y, header.Filename),
		},
	)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/form", handler).Methods(http.MethodPost)

	log.Fatal(
		http.ListenAndServe(":3000", r),
	)
}

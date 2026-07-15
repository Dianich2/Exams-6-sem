package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	segments := strings.Split(r.URL.Path, "/")
	if segments[1] == "sum" {
		x, err := strconv.Atoi(r.URL.Query().Get("x"))
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(
				Response{
					Code:    400,
					Message: "Bad Request",
				},
			)
			return
		}
		y, err := strconv.Atoi(r.URL.Query().Get("y"))
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(
				Response{
					Code:    400,
					Message: "Bad Request",
				},
			)
			return
		}
		json.NewEncoder(w).Encode(
			Response{
				Code:    200,
				Message: fmt.Sprintf("%d + %d = %d", x, y, x+y),
			},
		)
		return
	} else {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(
			Response{
				Code:    404,
				Message: "Not Found",
			},
		)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sum", handler).Methods(http.MethodGet)

	log.Fatal(
		http.ListenAndServe(":3000", r),
	)
}

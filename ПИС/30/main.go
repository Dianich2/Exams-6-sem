package main

import (
	"encoding/json"
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
	filename := mux.Vars(r)["filename"]
	inf, err := os.Stat(filename)
	if err == nil {
		if !inf.IsDir() {
			w.Header().Set("Content-Disposition", "attachment; filename="+filename)
			w.Header().Set("Content-Type", "application/octet-stream")
			http.ServeFile(w, r, filename)
			return
		} else {
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(
				Response{
					Code:    400,
					Message: "Bad Request: not a file",
				},
			)
			return
		}
	} else if os.IsNotExist(err) {
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			Response{
				Code:    404,
				Message: "Not Found: file not found",
			},
		)
		return
	} else {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			Response{
				Code:    500,
				Message: "Server error",
			},
		)
		return
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/download/{filename}", handler).Methods(http.MethodGet)

	log.Fatal(
		http.ListenAndServe(":3000", r),
	)
}

package main

import (
	"fmt"
	"net/http"
)

var validToken = "token456"

func profile(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token != "Bearer "+validToken {
		http.Error(w, "unauthorized", 401)
		return
	}

	fmt.Fprintln(w, "User Profile")
}

func main() {
	http.HandleFunc("/profile", profile)

	fmt.Println("Resource Server :8082")
	http.ListenAndServe(":8082", nil)
}

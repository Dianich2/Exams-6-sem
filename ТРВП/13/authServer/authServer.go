package main

import (
	"fmt"
	"net/http"
)

var authCode = "abc123"
var accessToken = "token456"

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login page")
}

func authorize(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	redirect := r.URL.Query().Get("redirect_uri")

	if clientID == "" || redirect == "" {
		http.Error(w, "invalid request", 400)
		return
	}

	http.Redirect(w, r, redirect+"?code="+authCode, 302)
}

func token(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	if code != authCode {
		http.Error(w, "invalid code", 401)
		return
	}

	fmt.Fprint(w, accessToken)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/authorize", authorize)
	http.HandleFunc("/token", token)

	fmt.Println("Auth Server :8083")
	http.ListenAndServe(":8083", nil)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var oauthState = "test-oauth"

var ConfigForOAuth = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:3000/callback",
	Scopes:       []string{"read:user"},
	Endpoint:     github.Endpoint,
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func rootHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(w, `<h1>OAuth2 test</h1>`)
	fmt.Fprintln(w, `<a href="/login">Login with GitHub</a>`)
}

func loginHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	url := ConfigForOAuth.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callbackHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	state := r.URL.Query().Get("state")
	if state != oauthState {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				Code:    http.StatusBadRequest,
				Message: "Bad Request: invalid state",
			},
		)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				Code:    http.StatusBadRequest,
				Message: "Bad Request: code is empty",
			},
		)
		return
	}

	token, err := ConfigForOAuth.Exchange(
		context.Background(),
		code,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Code:    http.StatusInternalServerError,
				Message: "Server error: cannot exchange code for token",
			},
		)
		return
	}

	client := ConfigForOAuth.Client(
		context.Background(),
		token,
	)

	resp, err := client.Get("https://api.github.com/user")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Code:    http.StatusInternalServerError,
				Message: "Server error: cannot get GitHub user",
			},
		)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Code:    http.StatusInternalServerError,
				Message: "Server error: cannot read GitHub response",
			},
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler).Methods(http.MethodGet)
	r.HandleFunc("/login", loginHandler).Methods(http.MethodGet)
	r.HandleFunc("/callback", callbackHandler).Methods(http.MethodGet)

	log.Printf("Server start on 3000\n")

	log.Fatal(
		http.ListenAndServe(":3000", r),
	)
}

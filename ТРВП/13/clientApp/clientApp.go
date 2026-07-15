package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8083/authorize?client_id=client123&redirect_uri=http://localhost:8081/callback"
	http.Redirect(w, r, url, 302)
}

func callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		http.Error(w, "no code", 400)
		return
	}

	resp, err := http.PostForm("http://localhost:8083/token",
		url.Values{
			"code": {code},
		},
	)

	if err != nil {
		http.Error(w, "token request failed: "+err.Error(), 500)
		return
	}

	defer resp.Body.Close()

	tokenBytes, _ := io.ReadAll(resp.Body)

	token := string(tokenBytes)
	token = strings.TrimSpace(token)
	fmt.Printf("TOKEN RAW: %q\n", token)

	req, _ := http.NewRequest("GET", "http://localhost:8082/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "api error: "+err.Error(), 500)
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	w.Write(body)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/callback", callback)

	println("Client App :8081")
	http.ListenAndServe(":8081", nil)
}

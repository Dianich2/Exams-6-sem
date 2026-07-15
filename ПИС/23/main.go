package main

import (
	"fmt"
	"log"
	"net/http"
)

func testHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	log.Printf(
		"Method=%s Path=%s",
		r.Method,
		r.URL.Path,
	)
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "POST":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "PUT":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "DELETE":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func userHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	log.Printf(
		"Method=%s Path=%s",
		r.Method,
		r.URL.Path,
	)
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "POST":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "PUT":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "DELETE":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	log.Printf(
		"Method=%s Path=%s",
		r.Method,
		r.URL.Path,
	)
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "POST":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "PUT":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	case "DELETE":
		fmt.Fprintf(w, "Method: %s, path: %s\n", r.Method, r.URL.Path)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tests", testHandler)
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/", handler)
	log.Printf("Server start on 3000\n")
	log.Fatal(
		http.ListenAndServe(":3000", nil),
	)
}

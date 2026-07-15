package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("BACK-GO")
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL.String())
	log.Println("Host:", r.Host)
	log.Println("RemoteAddr:", r.RemoteAddr)
	log.Println("X-Real-IP:", r.Header.Get("X-Real-IP"))
	log.Println("X-Forwarded-For:", r.Header.Get("X-Forwarded-For"))
	log.Println("X-Forwarded-Proto:", r.Header.Get("X-Forwarded-Proto"))

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "Hello from Go backend!")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Запрос пришел на back-go.")
	fmt.Fprintln(w, "Это внутреннее приложение, которое скрыто за Nginx reverse proxy.")
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Server started on port 3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Server   string    `json:"server"`
	Hostname string    `json:"hostname"`
	Time     time.Time `json:"time"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	hostname, _ := os.Hostname()

	serviceName := os.Getenv("SERVICE_NAME")

	resp := Response{
		Server:   serviceName,
		Hostname: hostname,
		Time:     time.Now(),
	}

	log.Printf("Request received by %s", serviceName)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func main() {

	http.HandleFunc("/inf", handler)

	log.Println("Server started")

	http.ListenAndServe(":3000", nil)
}

package main

import (
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(host))
}

func main() {
	log.Println("Starting http server...")
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

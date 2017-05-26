package main

import (
	"net/http"
	"log"
)

func sendOK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x-brady-header", "coolguy")
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", sendOK) // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
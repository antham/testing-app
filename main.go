package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const port = ":8080"

type Response struct {
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(Response{Message: "Hello World !!!!"})
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	log.Printf("Starting server on %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

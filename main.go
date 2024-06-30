package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World</h1>"))
		log.Println("Ok!")
	})
	http.ListenAndServe(":8080", nil)
}

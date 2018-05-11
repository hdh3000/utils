package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/scratch.html", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving scratch.html")
		http.ServeFile(w, r, "/Users/hdh/src/utils/go/src/utils/main/checklist/assets/scratch.html")
	}))

	http.Handle("/scratch.css", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving scratch.css")
		http.ServeFile(w, r, "/Users/hdh/src/utils/go/src/utils/main/checklist/assets/scratch.css")
	}))

	if err := http.ListenAndServe("127.0.0.1:3030", http.DefaultServeMux); err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}

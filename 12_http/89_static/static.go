package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs) // Handle registers the handler for the given pattern (/ in this case)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil)) // localhost:3000
}

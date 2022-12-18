package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8081

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

type application struct {
	Domain string
}

func main() {
	// set application config
	var app application

	// read from command line

	// connect to the database

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	http.HandleFunc("/", Hello)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, Index page!")
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, About page!")
}

func contactHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, Contact page!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

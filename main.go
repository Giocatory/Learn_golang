package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Guestbook struct {
	SignatureCount int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := "signatures"
	html, err := template.ParseFiles("./template/view.html")
	check(err)

	guestbook := Guestbook{
		SignatureCount: len(signatures),
	}

	err = html.Execute(writer, guestbook)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}

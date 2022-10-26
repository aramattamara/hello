package main

import (
	"io"
	"log"
	"net/http"
)

func main7() {
	callhttp()
}

func callhttp() {
	resp, err := http.Get("https://pkg.go.dev/net/http#pkg-overview")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func sendsqldata() {}

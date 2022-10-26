package main

import (
	"io"
	"log"
	"net/http"
)

//	func main() {
//		fmt.Println("Hello, World!")
//	}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main3() {
	x, y := split(17)
	log.Printf("%v, %v", x, y)
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

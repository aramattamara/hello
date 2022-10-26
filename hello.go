package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/mattn/go-sqlite3"
)

//	func main() {
//		fmt.Println("Hello, World!")
//	}

func main() {
	//callhttp()
	//dbconnect()
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

func dbconnect() {
	// The `sql.Open` function opens a new `*sql.DB` instance. We specify the driver name
	// and the URI for our database. Here, we're using a Postgres URI
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")

	res, err := db.Exec("INSERT INTO test VALUES(?)", "gbhyv")
	if err != nil {
		log.Fatalf("unable to insert data in database: %v", err)
	}
	fmt.Printf("insert is succesful %v", res)
}

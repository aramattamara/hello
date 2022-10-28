package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/mattn/go-sqlite3"
)

//	func main() {
//		fmt.Println("Hello, World!")
//	}

func main() {
	result := callhttp()
	dbconnect(result)
}

func callhttp() string {
	token := read_token()
	url := "https://api.telegram.org/bot" + token + "/getUpdates"

	resp, err := http.Get(url)
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
	return sb
}

func dbconnect(sb string) {
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

	res, err := db.Exec("INSERT INTO test VALUES(?)", sb)
	if err != nil {
		log.Fatalf("unable to insert data in database: %v", err)
	}
	fmt.Printf("insert is succesful %v", res)
}

func read_token() string {
	readFile, err := os.Open("token.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		token := fileScanner.Text()
		return token
	}
	return ""
}

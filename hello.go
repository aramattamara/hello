package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
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
	var result string = callHttp()
	log.Printf("received: %v\n\n", result)

	parsed, err := parseJson(result)
	if err != nil {
		fmt.Printf("error while parsing json %v\n", err)
		return
	}
	log.Printf("res: %+v\n\n", parsed)

	for _, v := range parsed.Result {
		message := v.Message

		if message == nil {
			message = v.EditedMessage
		}

		var messageId float64
		messageId = message.MessageId

		var messageJson []byte
		messageJson, err = json.Marshal(message)
		if err != nil {
			fmt.Printf("can't marshall json %v\n", err)
			return
		}
		fmt.Printf("messageJson: %v\n\n", string(messageJson))

		var fromUsername string
		fromUsername = message.From.Username

		saveToSqlite(
			int32(messageId),
			string(messageJson),
			(fromUsername),
		)
	}
}

func callHttp() string {
	token, err := readToken()
	if err != nil {
		log.Panicf("Cannot get token: %v", err)
	}
	url := "https://api.telegram.org/bot" + *token + "/getUpdates"

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
	return sb
}

func saveToSqlite(messageId int32, sb string, fromUsername string) {
	// The `sql.Open` function opens a new `*sql.DB` instance. We specify the driver name
	// and the URI for our database. Here, we're using a Postgres URI
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	res, err := db.Exec("INSERT OR REPLACE INTO message(id, from_username, content) VALUES(?,?,?)",
		messageId, fromUsername, sb)
	if err != nil {
		log.Fatalf("unable to insert data in database: %v", err)
	}
	lastInsertId, err := res.LastInsertId()
	rowsAffected, err := res.RowsAffected()
	fmt.Printf("insert is succesful, LastInsertId: %v, RowsAffected: %v \n\n",
		lastInsertId, rowsAffected)
}

func readToken() (*string, error) {
	fileName := "token.txt"
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		token := fileScanner.Text()
		return &token, nil
	}
	err = fmt.Errorf("nothing in file %v", fileName)
	return nil, err
}

func parseJson(sb string) (*Response, error) {
	bb := []byte(sb)

	response := &Response{}

	err := json.Unmarshal(bb, response)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

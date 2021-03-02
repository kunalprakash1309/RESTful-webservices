package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Book holds data of a book
type Book struct {
	ID int
	ISBN string
	Author string
	PublishedYear string
}

func main() {
	// File open for reading, writing and appending
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}
	defer f.Close()
	
	// This attaches program logs to file
	log.SetOutput(f)

	// Function handler for handling requests
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
		log.Printf("%q", r.UserAgent())

		book := Book{
			ID: 123,
			ISBN: "0-201-03801-3",
			Author: "Donald Knuth",
			PublishedYear: "1968",
		}
		// convert struct to JSON using marshal
		jsonData, err := json.Marshal(book)
		if err != nil {
			fmt.Println("error in marshiling the struct to json")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	s := &http.Server{
		Addr: ":8000",
	}
	log.Fatalln(s.ListenAndServe())
}
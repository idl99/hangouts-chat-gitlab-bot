package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/idl99/hangouts-chat-gitlab-bot/handlers"
)

var logger = log.New(os.Stderr, "logger: ", log.Lshortfile)
var port = ":8080"

func main() {
	http.HandleFunc("/", handlers.GitlabHandler)
	fmt.Println("Listening to requests on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

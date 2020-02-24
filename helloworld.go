package main

import (
	"fmt"
	"net/http"
	"os"
)
import _ "net/http"

func main() {
	fmt.Println("--------------------- Starting ")

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	if err := http.ListenAndServe(port(), nil); err != nil {
		fmt.Print(err)
		return
	}
}

func echo(writer http.ResponseWriter, request *http.Request) {

	s := request.URL.Query()["message"][0]

	writer.Header().Add("Accept","text/plain")
	writer.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(writer, s); err != nil {
		return
	}
}

func port() string {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		return ":8000"
	}
	return ":" + port
}

func index(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(writer, "kedise"); err != nil {
		return
	}

}

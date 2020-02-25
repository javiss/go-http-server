package api

import (
	"books/book.go"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("--------------------- Starting ")

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	http.HandleFunc("/api/books", Book)

	if err := http.ListenAndServe(port(), nil); err != nil {
		fmt.Print(err)
		return
	}
}

func echo(writer http.ResponseWriter, request *http.Request) {

	s := request.URL.Query()["message"][0]

	writer.Header().Add("Accept", "text/plain")
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

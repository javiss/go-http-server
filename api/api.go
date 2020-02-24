package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Book struct {
	Title string
	Author string
	ISBN string
}

func ToJson(book Book) []byte {
	ToJson,err := json.Marshal(book)
	if err != nil{
		panic(err)
	}
	return ToJson
}



func main() {
	fmt.Println("--------------------- Starting ")

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	if err := http.ListenAndServe(port(), nil); err != nil {
		fmt.Print(err)
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

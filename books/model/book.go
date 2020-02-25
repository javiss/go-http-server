package model

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) ToJson(book Book) []byte {
	ToJson, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	return ToJson
}

func (b Book) FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

func BooksHandleFunc(writer http.ResponseWriter, r *http.Request) {
}

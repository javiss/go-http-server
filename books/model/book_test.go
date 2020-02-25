package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookFromJSON(t *testing.T) {

	book := Book{}
	input := []byte(`{"Title":"book0","Author":"Auth0","ISBN":"ISBN0"}`)

	bookExpected := Book{Title: "book0", Author: "Auth0", ISBN: "ISBN0"}
	book = book.FromJSON(input)
	assert.Equal(t, bookExpected, book, "bad")

}

func TestBookToJSON(t *testing.T) {

	book := Book{Title: "book0", Author: "Auth0", ISBN: "ISBN0"}

	json := book.ToJson(book)
	assert.Equal(t, `{"title":"book0","author":"Auth0","isbn":"ISBN0"}`, string(json), "bad")

}

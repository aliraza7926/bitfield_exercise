package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "For the Love of Go",
		Author: "John Arundel",
		Copies: 3,
	}
}

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

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "For the Love of Go",
		Author: "John Arundel",
		Copies: 3,
	}

	want := 2
	result := bookstore.Buy(b)
	got := result.Copies

	if want != got {
		t.Errorf("want %d copies got %d", want, got)
	}

}

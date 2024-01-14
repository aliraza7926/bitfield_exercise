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
	result, err := bookstore.Buy(b)
	got := result.Copies

	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %d copies got %d", want, got)
	}

}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

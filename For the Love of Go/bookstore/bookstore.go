package bookstore

import (
	"errors"
)

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, ID int) Book {
	for _, book := range catalog {
		if book.ID == ID {
			return book
		}
	}
	return Book{}
}

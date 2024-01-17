package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book.
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

// Catalog is a map represents information about all book
type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	var result = []Book{}
	for _, book := range c {
		result = append(result, book)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	b, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving

}

func (b *Book) SetPriceCents(price int) error {
	b.PriceCents = price
	return nil
}

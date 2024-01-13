package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) Book {
	b.Copies--
	return b
}

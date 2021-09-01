package bookstore

import "fmt"

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID int
	PriceCents int
	DiscountPercent int
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("bad price %d (must not be negative)", price)
	}
	b.PriceCents = price
	return nil
}
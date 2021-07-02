package model

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int64
	CategoryID  uint
	Category    *Category
	Image       string
	IsSale      bool
}

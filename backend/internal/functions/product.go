package functions

import "errors"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func ValidateProduct(p Product) error {
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}
	if p.Price <= 0 {
		return errors.New("price must be positive")
	}
	if p.Stock < 0 {
		return errors.New("stock cannot be negative")
	}
	return nil
}

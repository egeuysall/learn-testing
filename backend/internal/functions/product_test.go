package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateProduct(t *testing.T) {
	tests := []struct {
		name      string
		product   Product
		wantError bool
		errorMsg  string
	}{
		{
			name:      "valid product",
			product:   Product{Name: "Laptop", Price: 999.99, Stock: 10},
			wantError: false,
		},
		{
			name:      "empty name",
			product:   Product{Name: "", Price: 999.99, Stock: 10},
			wantError: true,
			errorMsg:  "name cannot be empty",
		},
		{
			name:      "negative price",
			product:   Product{Name: "Charger", Price: -120.99, Stock: 10},
			wantError: true,
			errorMsg:  "price must be positive",
		},
		{
			name:      "negative stock",
			product:   Product{Name: "Charger", Price: 120.99, Stock: -10},
			wantError: true,
			errorMsg:  "stock cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProduct(tt.product)

			if tt.wantError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

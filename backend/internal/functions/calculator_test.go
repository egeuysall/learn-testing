package functions

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		wantError bool
	}{
		{"normal division", 10, 2, 5, false},
		{"negative division", -10, 2, -5, false},
		{"divide by zero", 10, 0, 0, true},
		{"dividing decimal numbers", 7, 2, 3.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("got %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

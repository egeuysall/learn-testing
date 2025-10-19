package functions

import (
	"testing"
)

func TestIsValidUsername(t *testing.T) {
	// Test with a valid username containing letters and numbers
	result := IsValidUsername("john123")
	if result != true {
		t.Errorf("IsValidUsername('john123') = %v; want true", result)
	}

	// Test with a username that is too short
	result = IsValidUsername("ab")
	if result != false {
		t.Errorf("IsValidUsername('ab') = %v; want false", result)
	}

	// Test with a username that is too long
	result = IsValidUsername("agshshajsjfkajdiorlpq")
	if result != false {
		t.Errorf("IsValidUsername('agshshajsjfkajdiorlpq') = %v; want false", result)
	}

	// Test with a username containing invalid characters (special characters)
	result = IsValidUsername("user@name")

	if result != false {
		t.Errorf("IsValidUsername('user@name') = %v; want false", result)
	}
}

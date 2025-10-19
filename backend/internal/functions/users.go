package functions

func IsValidUsername(username string) bool {
	length := len(username)
	if length < 3 || length > 20 {
		return false
	}

	for _, char := range username {
		isLetter := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
		isNumber := char >= '0' && char <= '9'

		if !isLetter && !isNumber {
			return false
		}
	}

	return true
}

package asciiartfs

import "errors"

// Check for valid of characters by runes from 32 to 126
func IsValid(s string) error {
	if len(s) > 400 {
		return errors.New("input text is too large")
	}
	for _, ch := range s {
		if (ch < ' ' && ch != '\n' && ch != '\r') || ch > '~' {
			return errors.New("invalid character in the input text")
		}
	}
	return nil
}

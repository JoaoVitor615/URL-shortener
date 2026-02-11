package encoder

import (
	"net/http"
	"strings"

	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
)

var (
	ErrInvalidEncodedString = apperrors.New("Invalid encoded string", http.StatusBadRequest)
)

// Encode converts a number to a base62 string
func Encode(num int) string {
	if num == 0 {
		return string(alphabet[0])
	}

	encoded := ""
	for num > 0 {
		encoded = string(alphabet[num%Base]) + encoded
		num = num / Base
	}
	return encoded
}

// Decode converts a base62 string to a number
func Decode(encoded string) (int, error) {
	encodedChars := strings.Split(encoded, "")
	decoded := 0
	for i := 0; i < len(encoded); i++ {
		if !ValidateCharacter(encodedChars[i]) {
			return 0, ErrInvalidEncodedString
		}
		decoded = decoded*Base + strings.Index(alphabet, encodedChars[i])
	}
	return decoded, nil
}

func ValidateCharacter(char string) bool {
	return strings.Contains(alphabet, char)
}

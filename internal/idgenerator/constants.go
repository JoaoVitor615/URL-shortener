package idgenerator

import "github.com/JoaoVitor615/URL-shortener/internal/encoder"

const (
	// The minimum number for the encoded ID having minimum 5 characters (N^4)
	minID = encoder.Base ^ 4

	// The maximum number for the encoded ID having maximum 7 characters (N^7 - 1)
	maxID = encoder.Base ^ 7 - 1
)

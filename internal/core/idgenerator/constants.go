package idgenerator

import (
	"math"

	"github.com/JoaoVitor615/URL-shortener/internal/core/encoder"
)

// not defined as const because of the math.Pow function
//
// it was not defined hardcoded for the flexibility of changing the base and
// the number of minimum and maximum characters.
var (
	// The minimum number for the encoded ID having minimum 5 characters (N^4)
	minID = int(math.Pow(float64(encoder.Base), 4))

	// The maximum number for the encoded ID having maximum 7 characters (N^7 - 1)
	maxID = int(math.Pow(float64(encoder.Base), 7) - 1)
)

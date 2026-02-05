package idgenerator

import "math/rand/v2"

// GenerateID generates a random ID between minID and maxID
func GenerateID() int { return rand.IntN(maxID-minID) + minID }

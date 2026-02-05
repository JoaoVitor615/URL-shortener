package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, "0", Encode(0))
}

func TestEncode_LargeNumber(t *testing.T) {
	assert.Equal(t, "1", Encode(1))
	assert.Equal(t, "a", Encode(10))
	assert.Equal(t, "PNFQ", Encode(12345678))
}

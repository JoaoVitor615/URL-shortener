package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, "0", Encode(0))
}

func TestEncode_LargeNumber(t *testing.T) {
	assert.Equal(t, "1Z", Encode(123))
	assert.Equal(t, "a", Encode(10))
	assert.Equal(t, "PNFQ", Encode(12345678))
}

func TestDecode(t *testing.T) {
	assert.Equal(t, 123, Decode("1Z"))
	assert.Equal(t, 10, Decode("a"))
	assert.Equal(t, 1, Decode("1"))
	assert.Equal(t, 12345678, Decode("PNFQ"))
}

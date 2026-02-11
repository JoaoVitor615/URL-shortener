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
	num, err := Decode("1Z")
	assert.NoError(t, err)
	assert.Equal(t, 123, num)
	num, err = Decode("a")
	assert.NoError(t, err)
	assert.Equal(t, 10, num)
	num, err = Decode("1")
	assert.NoError(t, err)
	assert.Equal(t, 1, num)
	num, err = Decode("PNFQ")
	assert.NoError(t, err)
	assert.Equal(t, 12345678, num)
}

func TestDecode_InvalidString(t *testing.T) {
	num, err := Decode("1/")
	assert.Error(t, err)
	assert.Equal(t, 0, num)
}

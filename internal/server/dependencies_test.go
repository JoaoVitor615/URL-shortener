package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDependencies(t *testing.T) {
	deps := NewDependencies()

	assert.NotNil(t, deps)
	assert.NotNil(t, deps.URLNumericHandler)
	assert.NotNil(t, deps.URLRandomHandler)
}

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewURLNumericService(t *testing.T) {
	service := NewURLNumericService()

	assert.NotNil(t, service)
	assert.IsType(t, &URLNumericService{}, service)
}

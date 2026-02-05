package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewURLRandomService(t *testing.T) {
	service := NewURLRandomService()

	assert.NotNil(t, service)
	assert.IsType(t, &URLRandomService{}, service)
}

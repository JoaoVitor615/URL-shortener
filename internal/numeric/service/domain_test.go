package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/JoaoVitor615/URL-shortener/internal/domain"
)

// MockRepositoryForDomain is a mock for testing NewNumericService
type MockRepositoryForDomain struct {
	mock.Mock
}

func (m *MockRepositoryForDomain) SaveURL(ctx context.Context, url *domain.URL[int]) error {
	args := m.Called(ctx, url)
	return args.Error(0)
}

func (m *MockRepositoryForDomain) GetURL(ctx context.Context, id int) (*domain.URL[int], error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.URL[int]), args.Error(1)
}

func (m *MockRepositoryForDomain) GetLongURL(ctx context.Context, longURL string) (*domain.URL[int], error) {
	args := m.Called(ctx, longURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.URL[int]), args.Error(1)
}

func TestNewNumericService_ReturnsInterface(t *testing.T) {
	mockRepo := new(MockRepositoryForDomain)
	
	service := NewNumericService(mockRepo)

	assert.NotNil(t, service)
	
	// Verify it implements INumericService interface
	var _ INumericService = service
}

func TestNewNumericService_ReturnsCorrectType(t *testing.T) {
	mockRepo := new(MockRepositoryForDomain)
	
	service := NewNumericService(mockRepo)

	// Cast to concrete type to verify internal structure
	numericService, ok := service.(*NumericService)
	assert.True(t, ok)
	assert.NotNil(t, numericService)
}

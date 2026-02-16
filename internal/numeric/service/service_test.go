package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/JoaoVitor615/URL-shortener/internal/core/encoder"
	"github.com/JoaoVitor615/URL-shortener/internal/domain"
)

// MockRepository is a mock implementation of IRepositoryShortURL[int]
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) SaveURL(ctx context.Context, url *domain.URL[int]) error {
	args := m.Called(ctx, url)
	return args.Error(0)
}

func (m *MockRepository) GetURL(ctx context.Context, id int) (*domain.URL[int], error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.URL[int]), args.Error(1)
}

func (m *MockRepository) GetLongURL(ctx context.Context, longURL string) (*domain.URL[int], error) {
	args := m.Called(ctx, longURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.URL[int]), args.Error(1)
}

func TestNewNumericService(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewNumericService(mockRepo)

	assert.NotNil(t, service)
}

func TestNumericService_GetLongURL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewNumericService(mockRepo)

		expectedID := 12345678
		shortURL := encoder.Encode(expectedID)
		expectedURL := &domain.URL[int]{
			ID:      expectedID,
			LongURL: "https://google.com",
		}

		mockRepo.On("GetURL", mock.Anything, expectedID).Return(expectedURL, nil)

		result, err := service.GetLongURL(shortURL)

		assert.NoError(t, err)
		assert.Equal(t, expectedURL, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid short URL", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewNumericService(mockRepo)

		result, err := service.GetLongURL("!!!invalid!!!")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, ErrInvalidNumericID, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewNumericService(mockRepo)

		expectedID := 12345678
		shortURL := encoder.Encode(expectedID)
		repoErr := errors.New("database error")

		mockRepo.On("GetURL", mock.Anything, expectedID).Return(nil, repoErr)

		result, err := service.GetLongURL(shortURL)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, repoErr, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestNumericService_CreateShortURL(t *testing.T) {
	t.Run("success - new URL", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewNumericService(mockRepo)

		url := domain.NewURL(0, "https://google.com")
		emptyURL := &domain.URL[int]{ID: 0, LongURL: ""}

		mockRepo.On("GetLongURL", mock.Anything, "https://google.com").Return(emptyURL, nil)
		mockRepo.On("SaveURL", mock.Anything, mock.AnythingOfType("*domain.URL[int]")).Return(nil)

		result, err := service.CreateShortURL(url)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Contains(t, result, "http")
		mockRepo.AssertExpectations(t)
	})

	t.Run("success - existing URL returns cached", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewNumericService(mockRepo)

		url := domain.NewURL(0, "https://google.com")
		existingURL := &domain.URL[int]{ID: 12345678, LongURL: "https://google.com"}

		mockRepo.On("GetLongURL", mock.Anything, "https://google.com").Return(existingURL, nil)

		result, err := service.CreateShortURL(url)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Contains(t, result, encoder.Encode(existingURL.ID))
		mockRepo.AssertExpectations(t)
	})

	t.Run("error - empty long URL", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewNumericService(mockRepo)

		url := domain.NewURL(0, "")

		result, err := service.CreateShortURL(url)

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Equal(t, domain.ErrLongURLRequired, err)
	})

	t.Run("error - GetLongURL fails", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewNumericService(mockRepo)

		url := domain.NewURL(0, "https://google.com")
		repoErr := errors.New("database error")

		mockRepo.On("GetLongURL", mock.Anything, "https://google.com").Return(nil, repoErr)

		result, err := service.CreateShortURL(url)

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Equal(t, repoErr, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error - SaveURL fails", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewNumericService(mockRepo)

		url := domain.NewURL(0, "https://google.com")
		emptyURL := &domain.URL[int]{ID: 0, LongURL: ""}
		saveErr := errors.New("save failed")

		mockRepo.On("GetLongURL", mock.Anything, "https://google.com").Return(emptyURL, nil)
		mockRepo.On("SaveURL", mock.Anything, mock.AnythingOfType("*domain.URL[int]")).Return(saveErr)

		result, err := service.CreateShortURL(url)

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Equal(t, saveErr, err)
		mockRepo.AssertExpectations(t)
	})
}

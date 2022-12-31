package memory

import (
	"errors"
	guid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"go-rest-api-with-db/internal/domain"
	"go-rest-api-with-db/internal/repositories"
	"testing"
)

func TestAuthorMemoryRepository_GetList(t *testing.T) {
}

func TestAuthorMemoryRepository_GetById(t *testing.T) {

	// Arrange
	expected, err := domain.NewAuthor("Hasan")
	if err != nil {
		t.Errorf("Author creation error %v", err)
	}

	store := map[guid.UUID]*domain.Author{}
	store[expected.ID] = expected

	repo := AuthorMemoryRepository{authors: store}

	testCases := []struct {
		name        string
		id          guid.UUID
		expectedErr error
	}{
		{
			name:        "no author by id",
			id:          guid.NewV4(),
			expectedErr: repositories.ErrAuthorNotFound,
		},
		{
			name:        "author by id",
			id:          expected.ID,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			actual, err := repo.GetById(tc.id)

			// Assert
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil {
				assert.NotNil(t, actual)
				assert.Equal(t, *expected, actual)
			}
		})
	}
}

func TestAuthorMemoryRepository_Add(t *testing.T) {
}

func TestAuthorMemoryRepository_Update(t *testing.T) {
}

func TestAuthorMemoryRepository_Delete(t *testing.T) {
}

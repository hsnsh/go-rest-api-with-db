package memory

import (
	"fmt"
	guid "github.com/satori/go.uuid"
	"go-rest-api-with-db/internal/domain"
	"go-rest-api-with-db/internal/repositories"
	"sync"
)

type AuthorMemoryRepository struct {
	authors map[guid.UUID]*domain.Author
	sync.Mutex
}

func New() *AuthorMemoryRepository {
	return &AuthorMemoryRepository{
		authors: make(map[guid.UUID]*domain.Author),
	}
}

func (a AuthorMemoryRepository) GetList() ([]domain.Author, error) {
	var authors []domain.Author
	for _, v := range a.authors {
		authors = append(authors, *v)
	}

	return authors, nil
}

func (a AuthorMemoryRepository) GetById(id guid.UUID) (domain.Author, error) {
	if author, ok := a.authors[id]; ok {
		return *author, nil
	}

	return domain.Author{}, repositories.ErrAuthorNotFound
}

func (a AuthorMemoryRepository) Add(input *domain.Author) error {
	a.Lock()
	defer a.Unlock()

	if a.authors == nil {
		a.authors = make(map[guid.UUID]*domain.Author)
	}

	// Make sure author is already in repo
	if _, ok := a.authors[input.ID]; ok {
		return fmt.Errorf("author is already exist %w", repositories.ErrAuthorAlreadyExist)
	}

	a.authors[input.ID] = input

	return nil
}

func (a AuthorMemoryRepository) Update(input *domain.Author) error {
	a.Lock()
	defer a.Unlock()

	if _, ok := a.authors[input.ID]; !ok {
		return fmt.Errorf("author does not exist %w", repositories.ErrAuthorNotFound)
	}

	a.authors[input.ID] = input

	return nil
}

func (a AuthorMemoryRepository) Delete(id guid.UUID) error {
	a.Lock()
	defer a.Unlock()

	if _, ok := a.authors[id]; !ok {
		return fmt.Errorf("author does not exist %w", repositories.ErrAuthorNotFound)
	}
	delete(a.authors, id)

	return nil
}

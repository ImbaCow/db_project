package teststore

import (
	"github.com/ImbaCow/bd_project/internal/app/store"
)

type testStore struct {
	repositoryStorage *repositoryStorage
}

// New ...
func New() store.Store {
	return &testStore{
		repositoryStorage: newRepositoryStorage(),
	}
}

// GetRepositoryStorage ...
func (s *testStore) GetRepositoryStorage() store.RepositoryStorage {
	return s.repositoryStorage
}

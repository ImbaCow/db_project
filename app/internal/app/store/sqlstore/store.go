package sqlstore

import (
	"database/sql"

	"github.com/ImbaCow/bd_project/internal/app/store"

	_ "github.com/go-sql-driver/mysql" // ...
)

type sqlStore struct {
	db                *sql.DB
	repositoryStorage *repositoryStorage
}

// Open ...
func Open(databaseURL string) (*sqlStore, error) {
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &sqlStore{
		db: db,
	}, nil
}

// Close ...
func (s *sqlStore) Close() {
	s.db.Close()
}

// GetRepositoryStorage ...
func (s *sqlStore) GetRepositoryStorage() store.RepositoryStorage {
	if s.repositoryStorage == nil {
		s.repositoryStorage = newRepositoryStorage(s)
	}

	return s.repositoryStorage
}

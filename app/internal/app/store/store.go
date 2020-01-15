package store

// Store ...
type Store interface {
	GetRepositoryStorage() RepositoryStorage
}

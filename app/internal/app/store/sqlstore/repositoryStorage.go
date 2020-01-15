package sqlstore

import (
	"github.com/ImbaCow/bd_project/internal/app/store"
)

type repositoryStorage struct {
	store             *sqlStore
	userRepository    *userRepository
	channelRepository *channelRepository
}

func newRepositoryStorage(store *sqlStore) *repositoryStorage {
	return &repositoryStorage{
		store: store,
	}
}

// GetUserRepository ...
func (r *repositoryStorage) GetUserRepository() store.UserRepository {
	if r.userRepository == nil {
		r.userRepository = newUserRepository(r.store)
	}

	return r.userRepository
}

// GetChannelRepository ...
func (r *repositoryStorage) GetChannelRepository() store.ChannelRepository {
	if r.channelRepository == nil {
		r.channelRepository = newChannelRepository(r.store)
	}

	return r.channelRepository
}

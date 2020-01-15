package teststore

import (
	"github.com/ImbaCow/bd_project/internal/app/store"
)

type repositoryStorage struct {
	userRepository    *userRepository
	channelRepository *channelRepository
}

func newRepositoryStorage() *repositoryStorage {
	return &repositoryStorage{
		userRepository:    newUserRepository(),
		channelRepository: newChannelRepository(),
	}
}

// GetUserRepository ...
func (r *repositoryStorage) GetUserRepository() store.UserRepository {
	return r.userRepository
}

// GetChannelRepository ...
func (r *repositoryStorage) GetChannelRepository() store.ChannelRepository {
	return r.channelRepository
}

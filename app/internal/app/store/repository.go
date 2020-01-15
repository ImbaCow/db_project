package store

import (
	"github.com/ImbaCow/bd_project/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) (int, error)
	Find(int) (*model.User, error)
	FindByLogin(string) (*model.User, error)
}

// ChannelRepository ...
type ChannelRepository interface {
	FindAll() ([]*model.Channel, error)
}

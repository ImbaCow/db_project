package teststore

import (
	"github.com/ImbaCow/bd_project/internal/app/model"
)

type channelRepository struct {
	channels []*model.Channel
}

func newChannelRepository() *channelRepository {
	return &channelRepository{
		channels: make([]*model.Channel, 0),
	}
}

// FindAll ...
func (r *channelRepository) FindAll() ([]*model.Channel, error) {
	return r.channels, nil
}

package sqlstore

import (
	"github.com/ImbaCow/bd_project/internal/app/model"
)

type channelRepository struct {
	store *sqlStore
}

func newChannelRepository(store *sqlStore) *channelRepository {
	return &channelRepository{
		store: store,
	}
}

// FindAll ...
func (r *channelRepository) FindAll() ([]*model.Channel, error) {
	channels := make([]*model.Channel, 0)

	rows, err := r.store.db.Query("SELECT id, name, display_name FROM channel")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		channel := &model.Channel{}
		if err := rows.Scan(&channel.ID, &channel.Name, &channel.DisplayName); err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}

	return channels, nil
}

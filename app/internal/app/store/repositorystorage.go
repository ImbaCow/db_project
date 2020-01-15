package store

// RepositoryStorage ...
type RepositoryStorage interface {
	GetUserRepository() UserRepository
	GetChannelRepository() ChannelRepository
}

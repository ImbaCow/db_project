package dbproject

import (
	"github.com/ImbaCow/bd_project/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

// Start ...
func Start(config *Config) error {
	store, err := sqlstore.Open(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer store.Close()
	sessionStore := sessions.NewFilesystemStore(config.SessionPath, []byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	return srv.Start(config)
}

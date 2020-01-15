package sqlstore

import (
	"github.com/ImbaCow/bd_project/internal/app/model"
)

type userRepository struct {
	store *sqlStore
}

func newUserRepository(store *sqlStore) *userRepository {
	return &userRepository{
		store: store,
	}
}

// Create ...
func (r *userRepository) Create(u *model.User) (int, error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}

	if err := u.BeforeCreate(); err != nil {
		return 0, err
	}

	res, execErr := r.store.db.Exec(
		"INSERT INTO user (login, passwordHash) VALUES (?, ?)",
		u.Login,
		u.PasswordHash,
	)
	if execErr != nil {
		return 0, execErr
	}
	id, err := res.LastInsertId()
	return int(id), err
}

// Find ..
func (r *userRepository) Find(id int) (*model.User, error) {
	user := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, login, passwordHash FROM user WHERE id = ?",
		id,
	).Scan(
		&user.ID,
		&user.Login,
		&user.PasswordHash,
	); err != nil {
		return nil, err
	}
	return user, nil
}

// FindByLogin ...
func (r *userRepository) FindByLogin(login string) (*model.User, error) {
	user := model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, login, passwordHash FROM user WHERE login = ?",
		login,
	).Scan(
		&user.ID,
		&user.Login,
		&user.PasswordHash,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

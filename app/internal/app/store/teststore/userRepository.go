package teststore

import (
	"errors"

	"github.com/ImbaCow/bd_project/internal/app/model"
)

type userRepository struct {
	users []*model.User
}

func newUserRepository() *userRepository {
	return &userRepository{
		users: make([]*model.User, 0),
	}
}

//Create ...
func (r *userRepository) Create(u *model.User) (int, error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}

	if err := u.BeforeCreate(); err != nil {
		return 0, err
	}

	newUser := *u
	id := len(r.users)
	newUser.ID = id
	r.users = append(r.users, &newUser)
	return newUser.ID, nil
}

// Find ...
func (r *userRepository) Find(id int) (*model.User, error) {
	if id >= len(r.users) {
		return nil, errors.New("index out of range")
	}
	return r.users[id], nil
}

// FindByLogin ...
func (r *userRepository) FindByLogin(login string) (*model.User, error) {
	for i := range r.users {
		if r.users[i].Login == login {
			return r.users[i], nil
		}
	}

	return nil, errors.New("User not found")
}

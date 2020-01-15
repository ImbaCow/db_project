package model

import (
	"github.com/ImbaCow/bd_project/internal/validation"
	ozzoValidation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"-"`
}

// Validate ...
func (u *User) Validate() error {
	return ozzoValidation.ValidateStruct(u,
		ozzoValidation.Field(&u.Login, ozzoValidation.Required),
		ozzoValidation.Field(&u.Password, ozzoValidation.By(validation.RequiredIf(u.PasswordHash == "")), ozzoValidation.Length(6, 100)),
	)
}

// BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.PasswordHash = enc
	}

	return nil
}

// IsPasswordEqual ...
func (u *User) IsPasswordEqual(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) == nil
}

func encryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

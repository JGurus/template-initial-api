package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User .
type User struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//HashPassword .
func (u *User) HashPassword(password string) error {
	sb, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return err
	}
	u.Password = string(sb)
	return nil
}

//ValidatedPassword .
func (u *User) ValidatedPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u User) String() string {
	return fmt.Sprintf("id: %d, username: %s, email: %s, password: %s created_at: %v, updated_at: %v\n", u.ID, u.Username, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)
}

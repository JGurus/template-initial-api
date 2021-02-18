package services

import (
	"log"

	"github.com/JGurus/template-initial-api/models"
)

//User .
type User struct {
	db UserServices
}

//NewUser .
func NewUser(db UserServices) User {
	return User{db}
}

//Migrate .
func (s User) Migrate() error {
	err := s.db.Migrate()
	if err != nil {
		return nil
	}
	return nil
}

//Create .
func (s User) Create(u *models.User) error {
	user := u
	user.HashPassword(u.Password)
	err := s.db.Create(user)
	if err != nil {
		return nil
	}
	return nil
}

//Update .
func (s User) Update(id int, u *models.User) error {
	user := u
	err := user.HashPassword(u.Password)
	if err != nil {
		return nil
	}
	err = s.db.Update(id, user)
	if err != nil {
		return nil
	}
	return nil
}

//Delete .
func (s User) Delete(id int) error {
	err := s.db.Delete(id)
	if err != nil {
		return nil
	}
	return nil
}

//GetAll .
func (s User) GetAll() ([]models.User, error) {
	users := make([]models.User, 0)
	users, err := s.db.GetAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

//GetByID .
func (s User) GetByID(id int) (models.User, error) {
	user := models.User{}
	user, err := s.db.GetByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetByUsername .
func (s User) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	user, err := s.db.GetByUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetByEmail .
func (s User) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	user, err := s.db.GetByEmail(email)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

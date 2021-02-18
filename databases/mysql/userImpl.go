package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JGurus/template-initial-api/models"
)

const (
	migrateUser = `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		username VARCHAR(25) NOT NULL,
		email VARCHAR(25) NOT NULL,
		password VARCHAR(60) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	 )`
	createUser    = "INSERT INTO users (username, email, password) VALUES(?,?,?)"
	updateUser    = "UPDATE users SET username = ?, email = ?, password = ?, updated_at = now() WHERE id = ? AND deleted_at IS NULL"
	deleteUser    = "UPDATE users SET deleted_at = now() WHERE id = ?"
	getAll        = "SELECT * FROM users WHERE deleted_at IS NULL"
	getByID       = "SELECT * FROM users WHERE id = ?"
	getByUsername = "SELECT * FROM users WHERE username = ?"
	getByEmail    = "SELECT * FROM users WHERE email = ?"
)

//UserImpl .
type UserImpl struct {
	db *sql.DB
}

//NewUserImpl .
func NewUserImpl(db *sql.DB) *UserImpl {
	return &UserImpl{db}
}

//Migrate .
func (dao *UserImpl) Migrate() error {
	stmt, err := dao.db.Prepare(migrateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migración exitosa")
	return nil
}

//Create .
func (dao *UserImpl) Create(u *models.User) error {
	stmt, err := dao.db.Prepare(createUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Username, u.Email, u.Password)
	if err != nil {
		return err
	}
	fmt.Println("Usuario creado exitosamente")
	return nil
}

//Update .
func (dao *UserImpl) Update(id int, u *models.User) error {
	stmt, err := dao.db.Prepare(updateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Username, u.Email, u.Password, id)
	if err != nil {
		return err
	}
	fmt.Println("Usuario actualizado exitosamente")
	return nil
}

//Delete .
func (dao *UserImpl) Delete(id int) error {
	stmt, err := dao.db.Prepare(deleteUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	println("Usuario eliminado exitosamente")
	return nil
}

//GetAll .
func (dao *UserImpl) GetAll() ([]models.User, error) {
	users := make([]models.User, 0)
	stmt, err := dao.db.Prepare(getAll)
	if err != nil {
		return users, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	fmt.Println("Usuarios consultados exitosamente")
	return users, nil
}

//GetByID .
func (dao *UserImpl) GetByID(id int) (models.User, error) {
	user := models.User{}
	stmt, err := dao.db.Prepare(getByID)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetByUsername return a User by username
func (dao UserImpl) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	stm, err := dao.db.Prepare(getByUsername)
	if err != nil {
		return user, err
	}
	defer stm.Close()
	row := stm.QueryRow(username)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetByEmail return a User by email
func (dao UserImpl) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	stm, err := dao.db.Prepare(getByEmail)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer stm.Close()
	row := stm.QueryRow(email)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

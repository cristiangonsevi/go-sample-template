package repository

import (
	"database/sql"
	"log"

	"example.com/internal/api/model"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUser(id int) (model.User, error)
	AddUser(user model.User) error
	UpdateUser(id int, user model.User) error
	DeleteUser(id int) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

var db *sql.DB

func NewUserRepositoryImpl() *UserRepositoryImpl {
	var err error
	db, err = sql.Open("mysql", "admin:admin@tcp(localhost:3306)/sample")
	if err != nil {
		log.Fatal(err)
	}

	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) GetUsers() ([]model.User, error) {
	var users []model.User

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetUser(id int) (model.User, error) {
	var user model.User

	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Age)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) AddUser(user model.User) error {
	_, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) UpdateUser(id int, user model.User) error {
	_, err := db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", user.Name, user.Age, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) DeleteUser(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

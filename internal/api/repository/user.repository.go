package repository

import (
	"database/sql"
	"log"

	"example.com/internal/api/model"
  _ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
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

	// defer db.Close()

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

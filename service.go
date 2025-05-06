package main

import (
	"database/sql"
)

type UserService interface {
	CreateUser(name, email string) error
	GetAllUsers() ([]User, error)
}

type PGUserService struct {
	db *sql.DB
}

func NewPGUserService(db *sql.DB) *PGUserService {
	return &PGUserService{db: db}
}

func (s *PGUserService) CreateUser(name, email string) error {
	_, err := s.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	return err
}

func (s *PGUserService) GetAllUsers() ([]User, error) {
	rows, err := s.db.Query("SELECT id, name, email FROM users ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

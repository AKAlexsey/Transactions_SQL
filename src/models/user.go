package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) CreateUser() error {
	query := `
		INSERT INTO users (email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	db := GetDB()

	now := time.Now()
	return db.QueryRow(query, u.Email, u.Password, now, now).Scan(&u.ID)
}

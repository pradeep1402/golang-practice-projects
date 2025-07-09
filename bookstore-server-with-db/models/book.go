package models

import "time"

type Book struct {
	Id        int       `db:"id"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PostFormBook struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

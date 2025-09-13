package models

import "time"

type Recipe struct {
	ID          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	MakingTime  string    `db:"making_time" json:"making_time"`
	Serves      string    `db:"serves" json:"serves"`
	Ingredients string    `db:"ingredients" json:"ingredients"`
	Cost        int       `db:"cost" json:"cost"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

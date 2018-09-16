package model

import "time"

type User struct {
	Id        int        `json:"id"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Todos     []*Todo    `json:"todos"`
}

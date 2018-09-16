package model

import "github.com/azbshiri/common/db"

type Todo struct {
	db.Model
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
}

package repository

import (
	"github.com/go-pg/pg"
)

type TodoWrapper struct {
	DB *pg.DB
}

func Todo(s Server) *TodoWrapper {
	return &TodoWrapper{s.DB()}
}

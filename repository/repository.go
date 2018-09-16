package repository

import "github.com/go-pg/pg"

type Server interface {
	DB() *pg.DB
}

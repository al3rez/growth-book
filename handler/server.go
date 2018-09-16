package handler

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

type Server struct {
	db *pg.DB
	*echo.Echo
}

func NewServer(db *pg.DB) *Server {
	server := &Server{db, echo.New()}
	server.Routes()
	return server
}

func (s *Server) DB() *pg.DB {
	return s.db
}

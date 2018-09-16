package main

import (
	"github.com/azbshiri/growth-book/config"
	"github.com/azbshiri/growth-book/handler"
	"github.com/go-pg/pg"
)

func main() {
	db := pg.Connect(&pg.Options{
		User:     config.DBUser,
		Password: config.DBPassword,
		Database: config.DBName,
	})

	server := handler.NewServer(db)
	server.Logger.Fatal(server.Start(":4000"))
}

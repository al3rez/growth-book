package main

import (
	"os"

	"github.com/azbshiri/growth-book/handler"
	"github.com/go-pg/pg"
)

func main() {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	server := handler.NewServer(db)
	server.Logger.Fatal(server.Start(":4000"))
}

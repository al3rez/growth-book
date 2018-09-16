package repository

import (
	"github.com/azbshiri/growth-book/model"
	"github.com/go-pg/pg"
)

type UserWrapper struct {
	DB *pg.DB
}

func User(s Server) *UserWrapper {
	return &UserWrapper{s.DB()}
}

func (t *UserWrapper) Create(email string) (*model.User, error) {
	user := &model.User{Email: email}
	err := t.DB.Insert(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (t *UserWrapper) Find(id int) (*model.User, error) {
	user := &model.User{Id: id}
	err := t.DB.Model(user).Column("user.*", "Todos").Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

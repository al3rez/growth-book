package handler

import (
	"net/http"
	"strconv"

	"github.com/azbshiri/growth-book/model"
	"github.com/azbshiri/growth-book/repository"
	"github.com/labstack/echo"
)

func (s *Server) createUser(c echo.Context) (err error) {
	user := new(model.User)
	err = c.Bind(user)
	if err != nil {
		return err
	}

	user, err = repository.User(s).Create(user.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) getUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := repository.User(s).Find(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

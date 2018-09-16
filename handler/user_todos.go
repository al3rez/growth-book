package handler

import (
	"net/http"
	"strconv"

	"github.com/azbshiri/growth-book/repository"
	"github.com/labstack/echo"
)

func (s *Server) getUserTodos(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := repository.User(s).Find(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user.Todos)
}

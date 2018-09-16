package handler

func (s *Server) Routes() {
	s.File("/", "public/index.html")

	s.POST("/users", s.createUser)
	s.GET("/users/:id", s.getUser)
	s.GET("/users/:id/todos", s.getUserTodos)
}

package management

import (
	m "bricking.com/mytreeholes/mongo"
	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo, ss m.Storage) {
	ep := &endpoint{
		s: ss,
	}

	e.GET("/", ep.hello)

	g := e.Group("/admin")
	g.GET("/login", ep.loginpage)
	g.POST("/login", ep.login)
}

type endpoint struct {
	s m.Storage
}

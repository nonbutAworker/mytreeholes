package management

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (e *endpoint) hello(c echo.Context) error {
	//cookie, _ := c.Cookie("username")
	//if err != nil {
	//	return c.String(http.StatusForbidden, "please login first")
	//}

	//log.Printf("cookie key: %s, cookie value: %s", cookie.Name, cookie.Value)
	return c.Render(http.StatusOK, "Home.html", "")
}

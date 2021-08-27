package management

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (e *endpoint) login(c echo.Context) error {
	v := new(basiclogin)
	if err := c.Bind(v); err != nil {
		return err
	}

	//validate password and username
	b := validteforpassword(v)

	if b {
		writeCookie(c)
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusForbidden)
	}

}

func (e *endpoint) loginpage(c echo.Context) error {
	return c.Render(http.StatusOK, "Login.html", "")
}

type basiclogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func validteforpassword(v *basiclogin) bool {
	if v.Name == "zz" && v.Password == "password" {
		return true
	} else {
		return false
	}
}

func writeCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	log.Println(cookie)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

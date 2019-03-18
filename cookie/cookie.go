// Lib for add ou remove cookie on Echo framework
package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Create cookie in the content on Echo framework
func Create(c echo.Context, name, value string, expires time.Time, domain string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires
	cookie.Domain = domain
	c.SetCookie(cookie)
}

// Delete cookie in the content on Echo framework
func Delete(c echo.Context, name string, expires time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
}

package cookie

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

const CookieName = "jwt"

func SetCookie(c echo.Context, token string) {
	// Cookieに保存
	cookie := new(http.Cookie)
	cookie.Domain = os.Getenv("FLONT_URL")
	// 	cookie.Domain = config.Config.FlontUrl
	cookie.Path = "/api"
	cookie.Name = CookieName
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}

func DeleteCookie(c echo.Context) error {
	// Cookieに保存
	cookie, err := c.Cookie(CookieName)
	if err != nil {
		return err
	}
	cookie.Domain = os.Getenv("FLONT_URL")
	// 	cookie.Domain = config.Config.FlontUrl
	cookie.Path = "/api"
	cookie.MaxAge = -1
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return nil
}

package jwt

import (
	"os"
	"time"

	"srb/domain/models"

	"github.com/labstack/echo/v4"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(c echo.Context, account models.Account) (string, error) {
	// ヘッダーとペイロードを設定
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    account.Uid,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	// ローカルで実行時に[Config.ini]を使用する場合。
	// token, err := claims.SignedString([]byte(config.Config.Secretkey))

	// Heroku用
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return token, err
}

func ReadToken(c echo.Context, cookie string) (string, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {

			// ローカルで実行時に[Config.ini]を使用する場合。
			// return []byte(config.Config.Secretkey), nil

			// Heroku用
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil

		})

	if err != nil {
		return "error", err
	}

	// トークンからクレームを取得
	claims := token.Claims.(*jwt.StandardClaims)
	err = nil
	return claims.Issuer, err
}

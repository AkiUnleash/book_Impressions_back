package controllers

import (
	"net/http"
	"srb/domain/database"
	"srb/domain/models"
	"srb/middle"
	"srb/repositories/result"
	"srb/until/cookie"
	"srb/until/jwt"
	"srb/until/uid"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register
// @Summary Register Account infomation in the database.
// @Description Use the account table.
// @tags account
// @Accept  json
// @Produce  json
// @Success 201 {string} string	"201 Created"
// @failure 409 {string} string	"409 It is already registered"
// @Router /api/account/signup [post]
func Register(c echo.Context) error {
	var data map[string]string

	if err := c.Bind(&data); err != nil {
		return err
	}

	var count int64
	database.DB.Model(&models.Account{}).Where("email = ?", data["email"]).Count(&count)
	if count > 0 {
		return result.Json(c, http.StatusConflict, "409 It is already registered.")
	}

	var account models.Account
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	account = models.Account{
		Username: data["name"],
		Uid:      uid.Generate(),
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&account)

	return result.Json(c, http.StatusCreated, "201 Created")
}

// Login
// @Summary If the infomation passed in the request body matches the data in the table, a cookie will be issued.
// @Description JWT certification
// @tags account
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"200 OK"
// @failure 400 {string} string	"400 incrrect password"
// @failure 401 {string} string	"401 unauthenticated"
// @failure 404 {string} string	"404 Not Found"
// @Router /api/account/login [post]
func Login(c echo.Context) error {

	// パラメータからデータを抽出
	var data map[string]string
	if err := c.Bind(&data); err != nil {
		return err
	}

	var account models.Account

	// アカウント登録済みを確認する
	database.DB.Where("email = ?", data["email"]).First(&account)
	if account.Id == 0 {
		return result.Json(c, http.StatusNotFound, "404 Not Found")
	}

	// パスワード不一致
	if err := bcrypt.CompareHashAndPassword(account.Password, []byte(data["password"])); err != nil {
		return result.Json(c, http.StatusBadRequest, "400 incrrect password")
	}

	// JSON Web Tokenの作成
	token, err := jwt.GenerateToken(c, account)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// Cookie保存
	cookie.SetCookie(c, token)

	// 結果出力
	return result.Json(c, http.StatusOK, "200 OK")
}

// Logout
// @Summary If the cookie exists, delete it.
// @Description JWT certification
// @tags account
// @Accept  json
// @Produce  json
// @Success 200
// @failure 401 {string} string	"401 unauthenticated"
// @Router /api/account/logout [post]
func Logout(c echo.Context) error {

	// Cookie削除
	if err := cookie.DeleteCookie(c); err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// 結果出力
	return result.Json(c, http.StatusOK, "200 OK")

}

// CurentUser
// @Summary Show infomation about the currently logged in user.
// @Description Browse Account table.
// @tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Account
// @Router /account/nowuser [get]
func CurrentUser(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// uidからUserを取得
	var account models.Account
	database.DB.Where("uid = ?", uid).First(&account)

	// JSONを返す
	return c.JSON(http.StatusOK, account)

}

// CurentUserUpdate
// @Summary Show infomation about the currently logged in user.
// @Description Browse Account table.
// @tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Account
// @failure 401 {string} string	"401 unauthenticated"
// @failure 404 {string} string	"404 Not Found"
// @Router /account/nowuser [get]
func CurrentUserUpdate(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// アカウントの存在確認
	var count int64
	database.DB.Model(&models.Account{}).Where("uid = ?", uid).Count(&count)
	if count == 0 {
		return result.Json(c, http.StatusNotFound, "404 Not Found")
	}

	// リクエストボディのデータを取得
	var data models.AccountUpdate
	if err := c.Bind(&data); err != nil {
		return err
	}

	// 更新処理
	database.DB.Model(&models.Account{}).Where("uid = ?", uid).Update(&data)

	// 結果出力
	return result.Json(c, http.StatusOK, "200 OK")

}

// CurentUserDelete
// @Summary Show infomation about the currently logged in user.
// @Description Browse Account table.
// @tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Account
// @failure 401 {string} string	"401 unauthenticated"
// @Router /account/nowuser [get]
func CurrentUserDelete(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// uidからUserを取得
	var account models.Account
	database.DB.Where("uid = ?", uid).Delete(&account)

	// Cookie削除
	if err := cookie.DeleteCookie(c); err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// 結果出力
	return result.Json(c, http.StatusOK, "200 OK")

}

package controllers

import (
	"net/http"
	"srb/domain/database"
	"srb/domain/models"
	"srb/middle"
	"srb/repositories/result"

	"github.com/labstack/echo/v4"
)

// ImpressionsWrite
// @Summary Diary registratinon process.
// @Description Can be executed only at login.
// @tags diary
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"200 OK"
// @failure 401 {string} string	"401 unauthenticated"
// @Router /diary [post]
func ImpressionsWrite(c echo.Context) error {

	// パラメータのBodyからデータをBind
	var data map[string]string
	if err := c.Bind(&data); err != nil {
		return err
	}

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// DBにデータ登録
	Impression := models.Impression{
		Uid:       uid,
		Isbn10:    data["isbn10"],
		Isbn13:    data["isbn13"],
		Booktitle: data["booktitle"],
		Imageurl:  data["imageurl"],
		Title:     data["title"],
		Body:      data["body"],
	}
	database.DB.Create(&Impression)

	return result.Json(c, http.StatusOK, "200 OK")

}

// Login
// @Summary Processing to display diary
// @Description Can be executed only at login.
// @tags diary
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Impression
// @Router /diary [Get]
func ImpressionsRead(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return err
	}

	var diary []models.Impression
	database.DB.Where("uid = ?", uid).Order("created_at DESC").Find(&diary)

	return c.JSON(http.StatusOK, diary)
}

// Login
// @Summary Process to delete diary.
// @Description Can be executed only at login.
// @tags diary
// @Accept  json
// @Produce  json
// @Success 200
// @Router /diary/:id [Delete]
func ImpressionsDelete(c echo.Context) error {
	id := c.Param("id")
	var diary []models.Impression
	database.DB.Where("id = ?", id).Delete(&diary)

	return c.JSON(http.StatusOK, "OK")
}

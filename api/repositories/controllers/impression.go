package controllers

import (
	"fmt"
	"net/http"
	"srb/domain/database"
	"srb/domain/models"
	"srb/middle"
	"srb/repositories/result"

	"github.com/labstack/echo/v4"
)

// ImpressionsWrite
// @Summary impression registration process.
// @Description Can be executed only at login.
// @tags impressions
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"200 OK"
// @failure 401 {string} string	"401 unauthenticated"
// @Router /api/impressions [post]
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
		Bookid:    data["bookid"],
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

// ImpressionsRead
// @Summary List of impressions
// @Description Can be executed only at login.
// @tags impressions
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Impression
// @failure 401 {string} string	"401 unauthenticated"
// @Router /api/impressions [Get]
func ImpressionsRead(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	var Impression []models.Impression
	database.DB.Where("uid = ?", uid).Order("created_at DESC").Find(&Impression)

	return c.JSON(http.StatusOK, Impression)
}

// ImpressionRead
// @Summary Display of impressions.
// @Description Can be executed only at login.
// @tags impression
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"200 OK"
// @failure 401 {string} string	"401 unauthenticated"
// @failure 404 {string} string	"404 Not Found"
// @Router /api/impression/:id [Get]
func ImpressionRead(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	id := c.Param("id")
	var Impression models.Impression
	database.DB.Where("uid = ?", uid).Where("id = ?", id).Find(&Impression)

	// 対象者データ無し
	if Impression.Id == 0 {
		return result.Json(c, http.StatusNotFound, "404 Not Found")
	}

	return c.JSON(http.StatusOK, Impression)
}

// ImpressionDelete
// @Summary Delete impressions.
// @Description Can be executed only at login.
// @tags impression
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"200 OK"
// @failure 401 {string} string	"401 unauthenticated"
// @failure 404 {string} string	"404 Not Found"
// @Router /api/impression/:id [Delete]
func ImpressionDelete(c echo.Context) error {
	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	id := c.Param("id")
	var Impression models.Impression
	database.DB.Where("uid = ?", uid).Where("id = ?", id).Find(&Impression)

	// 対象者データ無し
	if Impression.Id == 0 {
		return result.Json(c, http.StatusNotFound, "404 Not Found")
	}

	database.DB.Delete(&Impression)

	return result.Json(c, http.StatusOK, "200 OK")
}

// ImpressionUpdate
// @Summary Update of impressions.
// @Description Can be executed only at login.
// @tags impression
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"200 OK"
// @failure 401 {string} string	"401 unauthenticated"
// @failure 404 {string} string	"404 Not Found"
// @Router /api/impression/:id [put]
func ImpressionUpdate(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// パラメータのBodyからデータをBind
	var data models.Impression
	if err := c.Bind(&data); err != nil {
		return err
	}

	fmt.Println(uid, data)

	// DBにデータ登録
	Impression := models.Impression{
		Uid:       uid,
		Bookid:    data.Bookid,
		Isbn10:    data.Isbn10,
		Isbn13:    data.Isbn13,
		Booktitle: data.Booktitle,
		Imageurl:  data.Imageurl,
		Title:     data.Title,
		Body:      data.Body,
	}

	database.DB.Model(&models.Impression{}).Where("uid = ?", uid).Where("id = ?", data.Id).Update(&Impression)

	return result.Json(c, http.StatusOK, "200 OK")
}

// ImpressionsSearch
// @Summary Search for books using Bookid as a key.
// @Description Can be executed only at login.
// @tags impressions
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Impression
// @failure 401 {string} string	"401 unauthenticated"
// @failure 404 {string} string	"404 Not Found"
// @Router /api/impressions/search/:bookid [Get]
func ImpressionsSearch(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return result.Json(c, http.StatusUnauthorized, "401 unauthenticated")
	}

	// クエリパラメータの取得
	bookid := c.QueryParam("bookid")

	var Impression []models.Impression
	database.DB.Where("uid = ?", uid).Where("bookid = ?", bookid).Order("created_at DESC").Find(&Impression)

	// 対象者データ無し
	if len(Impression) == 0 {
		return result.Json(c, http.StatusNotFound, "404 Not Found")
	}

	return c.JSON(http.StatusOK, Impression)
}

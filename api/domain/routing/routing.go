package routing

import (
	"srb/repositories/controllers"

	_ "github.com/go-sql-driver/mysql"

	_ "srb/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routing() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
	}))

	e.POST("api/account/signup", controllers.Register)
	e.POST("api/account/login", controllers.Login)
	e.POST("api/account/logout", controllers.Logout)
	e.GET("api/account/nowuser", controllers.CurrentUser)
	e.PUT("api/account/nowuser", controllers.CurrentUserUpdate)
	e.DELETE("api/account/nowuser", controllers.CurrentUserDelete)

	e.GET("api/impressions", controllers.ImpressionsRead)
	e.POST("api/impressions", controllers.ImpressionsWrite)
	e.GET("api/impression/:id", controllers.ImpressionRead)
	e.PUT("api/impression/:id", controllers.ImpressionUpdate)
	e.DELETE("api/impression/:id", controllers.ImpressionDelete)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8082"))
}

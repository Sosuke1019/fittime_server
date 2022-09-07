package router

import (
	_ "net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouter(e *echo.Echo) error {

	// 諸々の設定(*1)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// api
	api := e.Group("/api")
	{
		// ping
		api.GET("/ping", GetPingHandler)

		// User
		apiUser := api.Group("/user")
		{
			apiUser.POST("/create", CreateUserHandler)


			apiUserId := apiUser.Group("/:userId")
			{
				apiUserId.PATCH("/profile", AddProfileHandler)

				apiUserId.POST("/menu", PostMenuHandler)
			}
			apiUser.PATCH("/:userId/profile", AddProfileHandler)
			apiUser.PATCH("/:userId/username", AddNameHandler)

		}

		// Search
		apiSearch := api.Group("/search")
		{
			apiSearch.GET("", SearchHandler)
		}

		// Auth
		apiAuth := api.Group("/auth")
		{
			apiAuth.POST("/login", LoginHandler)
			apiAuth.POST("/:userId/logout", LogoutHandler)
		}
	}

	// 8000番のポートを開く(*2)
	err := e.Start(":8000")
	return err
}

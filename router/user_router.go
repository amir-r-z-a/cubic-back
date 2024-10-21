package router

import (
	"github.com/amir-r-z-a/cubic-back/middleware"
	"github.com/amir-r-z-a/cubic-back/services"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(app *gin.Engine, userService *services.UserService) {

	app.Use(middleware.CORSMiddleware())

	v0Public := app.Group("api/v0/public")
	{
		v0Public.POST("/signup", userService.SignUp)
		v0Public.POST("/signin", userService.SignIn)
	}

	v0Private := app.Group("api/v0/private")
	v0Private.Use(middleware.AuthMiddleware(userService))
	{
		v0Private.GET("/user", userService.GetUser)
	}
}

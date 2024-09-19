package router

import (
	"github.com/amir-r-z-a/cubic-back/services"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(app *gin.Engine, userService *services.UserService) {
	app.POST("/signup", userService.SignUp)

}

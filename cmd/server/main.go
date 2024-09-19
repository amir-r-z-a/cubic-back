package main

import (
	"log/slog"
	"os"

	"github.com/amir-r-z-a/cubic-back/config"
	"github.com/amir-r-z-a/cubic-back/repos"
	"github.com/amir-r-z-a/cubic-back/router"
	"github.com/amir-r-z-a/cubic-back/server"
	"github.com/amir-r-z-a/cubic-back/services"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	appConfig := config.LoadConfig()

	appConfig.Logger = logger

	appRepo := repos.InitRepo(appConfig)

	userRepo := repos.NewUserRepo(appRepo)

	userService := 	services.NewUserService(userRepo, appConfig)
	

	app := server.NewWebServer()

	router.AddUserRoutes(app, userService)

	server.Run(app, appConfig)

	

	


}



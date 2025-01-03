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
	scoreRepo := repos.NewScoreRepository(appRepo)
	commentRepo := repos.NewCommentRepository(appRepo)

	userService := services.NewUserService(userRepo, appConfig, appConfig.SecretKey)
	scoreService := services.NewScoreService(scoreRepo, appConfig)
	commentService := services.NewCommentService(commentRepo, appConfig)

	app := server.NewWebServer()

	router.AddUserRoutes(app, userService)
	router.SetupScoreRoutes(app, scoreService, userService)
	router.SetupCommentRoutes(app, commentService, userService)

	server.Run(app, appConfig)
}

package app

import (
	"go-basic-poc/internal/config"
	"go-basic-poc/internal/routes"
	"go-basic-poc/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	utils.InitLogger()
	defer utils.Logger.Sync()

	err = utils.InitDB(cfg)
	if err != nil {
		utils.Logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	r := gin.Default() // Create a Gin router

	routes.SetupRoutes(r) // Pass the Gin router to routes

	utils.Logger.Info("Starting server", zap.String("port", cfg.Port))
	err = r.Run(":" + cfg.Port) // Start the Gin server
	if err != nil {
		utils.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}

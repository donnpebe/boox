package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/donnpebe/boox/pkg/handler"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	handlers := handler.NewHandler()

	app := fiber.New()
	app.Use(logger.New())
	handlers.RegisterRoutes(app)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("BOOX")
	viper.SetDefault("port", ":8880")

	go func() {
		if err := app.Listen(viper.GetString("port")); err != nil {
			logrus.Fatalf("failed to listen: %v", err)
		}
	}()

	logrus.Println("Application started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Shutting down...")
	if err := app.Shutdown(); err != nil {
		logrus.Errorf("error occurred on server shutting down: %v", err)
	}
}

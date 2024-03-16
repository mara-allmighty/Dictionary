package main

import (
	"dictionary/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	logger := log.New("dict")

	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()
	api := router.Group("api")

	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)

	router.Logger.Fatal(router.Start(":8000"))
}

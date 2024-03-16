package main

import (
	"dictionary/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	// создаем логгер
	logger := log.New("dict")

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// прописываем пути
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}

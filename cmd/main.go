package main

import (
	"dictionary/internal/service"
	"dictionary/pkg/logs"

	"github.com/labstack/echo/v4"
)

func main() {
	// создаем логгер
	logger := logs.NewLogger(false)

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()

	// создаем группу api
	api := router.Group("api")

	// пути для word
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.PUT("/word/:id", svc.UpdateWordById)
	api.DELETE("/word/:id", svc.DeleteWordById)

	// пути для report
	api.GET("/report/:id", svc.GetReport)
	api.POST("/reports", svc.CreateReport)
	api.PUT("/report/:id", svc.UpdateReport)
	api.DELETE("/report/:id", svc.DeleteReport)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}

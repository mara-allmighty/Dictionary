package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) SearchClosestWords(c echo.Context) error {
	title := c.QueryParam("title")

	data, err := s.wordsRepo.RSearchClosestWords(title)
	if err != nil {
		return c.JSON(s.NewError(InvalidParams))
	}
	return c.JSON(http.StatusOK, data)
}

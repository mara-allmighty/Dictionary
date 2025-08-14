package service

import (
	"github.com/labstack/echo/v4"
)

// localhost:8000/api/report/:id
func (s *Service) GetReport(c echo.Context) error {
	return nil
}

// localhost:8000/api/report/words
func (s *Service) CreateReport(c echo.Context) error {
	return nil
}

// localhost:8000/api/report/:id
type UpdateReport struct {
	Title       string `json:"title" validate:"required,min=1"`
	Translation string `json:"translation" validate:"required,min=1"`
}

func (s *Service) UpdateReport(c echo.Context) error {
	return nil
}

// localhost:8000/api/report/:id
func (s *Service) DeleteReport(c echo.Context) error {
	return nil
}

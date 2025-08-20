package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GET - localhost:8000/api/report/:id
func (s *Service) GetReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	report, err := s.reportsRepo.GetReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: report})
}

// POST - localhost:8000/api/reports
func (s *Service) CreateReport(c echo.Context) error {
	var report Report
	err := c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	err = s.reportsRepo.CreateReport(report.Title, report.Overview)

	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// UPDATE - localhost:8000/api/report/:id
func (s *Service) UpdateReport(c echo.Context) error {
	// парсим JSON из тела запроса
	var req Report
	if err := c.Bind(&req); err != nil {
		s.logger.Error("Invalid JSON:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body, expected JSON",
		})
	}

	// парсим id-шник из URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	// обновляем репорт в БД
	err = s.reportsRepo.UpdateReport(id, req.Title, req.Overview)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, Response{Object: "Ok"})
}

// DELETE - localhost:8000/api/report/:id
func (s *Service) DeleteReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	err = s.reportsRepo.DeleteReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, Response{Object: "Ok"})
}

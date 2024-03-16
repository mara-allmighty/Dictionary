package service

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

const (
	InvalidParams       = "invalid params"
	InternalServerError = "internal error"
)

type Service struct {
	db     *sql.DB
	logger echo.Logger
}

func NewService(db *sql.DB, logger echo.Logger) *Service {
	return &Service{
		db:     db,
		logger: logger,
	}
}

type Response struct {
	Object       any    `json:"object,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func (r *Response) Error() string {
	return r.ErrorMessage
}

func (s *Service) NewError(err string) error {
	return &Response{ErrorMessage: err}
}

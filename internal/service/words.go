package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetWordById ищем слово по id
// localhost:8000/api/word/:id
func (s *Service) GetWordById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	word, err := repo.RGetWordById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: word})
}

// CreateWords добавляем в базу новые слова в базу
// localhost:8000/api/words
func (s *Service) CreateWords(c echo.Context) error {
	var wordSlice []Word
	err := c.Bind(&wordSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	for _, word := range wordSlice {
		err = repo.CreateNewWords(word.Title, word.Translation)
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// UpdateWordById - редактируем слово в БД
// localhost:8000/api/word/:id
type UpdateData struct {
	Title       string `json:"title" validate:"required,min=1"`
	Translation string `json:"translation" validate:"required,min=1"`
}

func (s *Service) UpdateWordById(c echo.Context) error {
	// Парсим JSON из тела запроса
	var req UpdateData
	if err := c.Bind(&req); err != nil {
		s.logger.Error("Invalid JSON:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body, expected JSON",
		})
	}

	// Парсим id-шник из URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	// Обновляем слово в БД
	err = s.wordsRepo.RUpdateWordById(id, req.Title, req.Translation)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, Response{Object: "Ok"})
}

// DeleteWordById - удаляем запись слова в БД
// localhost:8000/api/word/:id
func (s *Service) DeleteWordById(c echo.Context) error {
	// Парсим id-шник из URL
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	// Удаляем слово из БД
	err = s.wordsRepo.RDeleteWordById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, Response{Object: "Ok"})
}

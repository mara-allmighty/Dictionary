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

func (s *Service) UpdateWordById(c echo.Context) error {
	// UpdateWordById - редактируем слово в БД
	// PUT_localhost:8000/api/word/:id

	title := c.FormValue("title")
	translation := c.FormValue("translation")
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	msg, err := repo.RUpdateWordById(id, title, translation)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, Response{Object: msg})
}

func (s *Service) DeleteWordById(c echo.Context) error {
	// DeleteWordById - удаляем запись слова в БД
	// DELETE_localhost:8000/api/word/:id

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	msg, err := repo.RDeleteWordById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, Response{Object: msg})
}

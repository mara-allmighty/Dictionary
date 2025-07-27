package words

import (
	"database/sql"
	"errors"
	"fmt"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// RGetWordById ищем слово по id
func (r *Repo) RGetWordById(id int) (*Word, error) {
	var word Word
	err := r.db.QueryRow(`SELECT id, title, translation FROM ru_en WHERE id = $1`, id).
		Scan(&word.Id, &word.Title, &word.Translation)
	if err != nil {
		return nil, err
	}

	return &word, nil
}

// CreateNewWords добавляет новые переводы в базу даных
func (r *Repo) CreateNewWords(word, translate string) error {
	_, err := r.db.Exec(`INSERT INTO ru_en (title, translation) VALUES ($1, $2)`, word, translate)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) RUpdateWordById(id int, updatedTitle, updatedTranslation string) (string, error) {
	// RUpdateWordById - редактирует запись в БД "dict"
	// Проверяет в БД существование записи по полученному id

	var word Word
	myErr := r.db.QueryRow(`SELECT id, title, translation FROM ru_en WHERE id = $1`, id).Scan(&word.Id, &word.Title, &word.Translation)
	if myErr == sql.ErrNoRows {
		return fmt.Sprintf("Word with ID %d doesn't exist", id), errors.New("WordNotFoundError")
	}

	_, err := r.db.Exec(
		`UPDATE ru_en
			SET title = $1,
				translation = $2
		WHERE id = $3`,
		updatedTitle,
		updatedTranslation,
		id,
	)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Word with ID %d updated successfully!", id)
	return msg, nil
}

func (r *Repo) RDeleteWordById(id int) (string, error) {
	// RDeleteWordById - удаляет запись в БД "dict"
	// Проверяет в БД существование записи по полученному id

	var word Word
	myErr := r.db.QueryRow(`SELECT id, title, translation FROM ru_en WHERE id = $1`, id).Scan(&word.Id, &word.Title, &word.Translation)
	if myErr == sql.ErrNoRows {
		return fmt.Sprintf("Word with ID %d doesn't exist", id), errors.New("WordNotFoundError")
	}

	_, err := r.db.Exec(`DELETE FROM ru_en WHERE id = $1`, id)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Word with ID %d was deleted successfully!", id)
	return msg, nil
}

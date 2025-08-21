package words

import (
	"database/sql"
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

// Ищет 100 похожих слов
func (r *Repo) RSearchClosestWords(title string) ([]Word, error) {
	// фильтруем базу
	rows, err := r.db.Query(`
        SELECT id, title, translation,
        FROM ru_en 
        WHERE similarity(title, $1) >= 0.2
        ORDER BY sim DESC
        LIMIT 100`,
		title,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Собираем 100 похожих слов
	var closestsWords []Word
	for rows.Next() {
		var closestWord Word
		err := rows.Scan(&closestWord.Id, &closestWord.Title, &closestWord.Translation)
		if err != nil {
			return nil, err
		}
		closestsWords = append(closestsWords, closestWord)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return closestsWords, nil
}

// CreateNewWords добавляет новые переводы в базу даных
func (r *Repo) CreateNewWords(word, translate string) error {
	_, err := r.db.Exec(`INSERT INTO ru_en (title, translation) VALUES ($1, $2)`, word, translate)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) RUpdateWordById(id int, updatedTitle, updatedTranslation string) error {
	var title string

	// Выполняем UPDATE и сразу получаем обновлённое значение
	err := r.db.QueryRow(
		`UPDATE ru_en
			SET title = $1, translation = $2
			WHERE id = $3
			RETURNING title`,
		updatedTitle, updatedTranslation, id,
	).Scan(&title)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("word with ID %d not found", id)
		}
		return fmt.Errorf("database error: %w", err)
	}

	return nil
}

func (r *Repo) RDeleteWordById(id int) error {
	var title string

	// Выполняем DELETE и сразу получаем удалённое значение
	err := r.db.QueryRow(
		`DELETE FROM ru_en
		 WHERE id = $1
		 RETURNING title`,
		id,
	).Scan(&title)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("word with ID %d not found", id)
		}
		return fmt.Errorf("database error: %w", err)
	}

	return nil
}

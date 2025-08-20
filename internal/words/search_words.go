package words

func (r *Repo) RSearchClosestWords(title string) ([]Word, error) {
	// фильтруем базу
	rows, err := r.db.Query(`
        SELECT id, title, translation,
               similarity(title, $1) as sim
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
		var similarity float64
		err := rows.Scan(&closestWord.Id, &closestWord.Title, &closestWord.Translation, &similarity)
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

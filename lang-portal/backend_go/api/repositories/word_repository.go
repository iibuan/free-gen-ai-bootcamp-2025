package repositories

import (
	"backend_go/api/models"
)

func GetWords(page, perPage int) ([]map[string]interface{}, int, error) {
	offset := (page - 1) * perPage
	rows, err := DB.Query(`
        SELECT w.id, w.bahasa_indonesia, w.english, 
               (SELECT COUNT(*) FROM word_review_items WHERE word_id = w.id AND correct = 1) AS correct_count,
               (SELECT COUNT(*) FROM word_review_items WHERE word_id = w.id AND correct = 0) AS wrong_count
        FROM words w
        LIMIT ? OFFSET ?
    `, perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var words []map[string]interface{}
	for rows.Next() {
		var word models.Word
		var correctCount, wrongCount int
		if err := rows.Scan(&word.ID, &word.BahasaIndonesia, &word.English, &correctCount, &wrongCount); err != nil {
			return nil, 0, err
		}
		wordMap := map[string]interface{}{
			"id":               word.ID,
			"bahasa_indonesia": word.BahasaIndonesia,
			"english":          word.English,
			"correct_count":    correctCount,
			"wrong_count":      wrongCount,
		}
		words = append(words, wordMap)
	}

	var totalItems int
	err = DB.QueryRow("SELECT COUNT(*) FROM words").Scan(&totalItems)
	if err != nil {
		return nil, 0, err
	}

	return words, totalItems, nil
}

func GetWord(id int) (map[string]interface{}, error) {
	var word models.Word
	var correctCount, wrongCount int
	err := DB.QueryRow(`
        SELECT w.id, w.bahasa_indonesia, w.english, 
               (SELECT COUNT(*) FROM word_review_items WHERE word_id = w.id AND correct = 1) AS correct_count,
               (SELECT COUNT(*) FROM word_review_items WHERE word_id = w.id AND correct = 0) AS wrong_count
        FROM words w
        WHERE w.id = ?
    `, id).Scan(&word.ID, &word.BahasaIndonesia, &word.English, &correctCount, &wrongCount)
	if err != nil {
		return nil, err
	}

	rows, err := DB.Query(`
        SELECT g.id, g.name
        FROM groups g
        JOIN words_groups wg ON g.id = wg.group_id
        WHERE wg.word_id = ?
    `, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	wordMap := map[string]interface{}{
		"id":               word.ID,
		"bahasa_indonesia": word.BahasaIndonesia,
		"english":          word.English,
		"correct_count":    correctCount,
		"wrong_count":      wrongCount,
		"groups":           groups,
	}

	return wordMap, nil
}

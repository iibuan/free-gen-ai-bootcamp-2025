package repositories

import (
	"backend_go/api/models"
	"time"
)

func GetGroups(page, perPage int) ([]map[string]interface{}, int, error) {
	offset := (page - 1) * perPage
	rows, err := DB.Query(`
        SELECT g.id, g.name, 
               (SELECT COUNT(*) FROM words_groups wg WHERE wg.group_id = g.id) AS word_count
        FROM groups g
        LIMIT ? OFFSET ?
    `, perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var groups []map[string]interface{}
	for rows.Next() {
		var group models.Group
		var wordCount int
		if err := rows.Scan(&group.ID, &group.Name, &wordCount); err != nil {
			return nil, 0, err
		}
		groupMap := map[string]interface{}{
			"id":         group.ID,
			"name":       group.Name,
			"word_count": wordCount,
		}
		groups = append(groups, groupMap)
	}

	var totalItems int
	err = DB.QueryRow("SELECT COUNT(*) FROM groups").Scan(&totalItems)
	if err != nil {
		return nil, 0, err
	}

	return groups, totalItems, nil
}

func GetGroup(id int) (map[string]interface{}, error) {
	var group models.Group
	var wordCount int
	err := DB.QueryRow(`
        SELECT g.id, g.name, 
               (SELECT COUNT(*) FROM words_groups wg WHERE wg.group_id = g.id) AS word_count
        FROM groups g
        WHERE g.id = ?
    `, id).Scan(&group.ID, &group.Name, &wordCount)
	if err != nil {
		return nil, err
	}

	groupMap := map[string]interface{}{
		"id":         group.ID,
		"name":       group.Name,
		"word_count": wordCount,
	}

	return groupMap, nil
}

func GetGroupWords(groupID, page, perPage int) ([]map[string]interface{}, int, error) {
	offset := (page - 1) * perPage
	rows, err := DB.Query(`
        SELECT w.id, w.bahasa_indonesia, w.english, 
               (SELECT COUNT(*) FROM word_review_items WHERE word_id = w.id AND correct = 1) AS correct_count,
               (SELECT COUNT(*) FROM word_review_items WHERE word_id = w.id AND correct = 0) AS wrong_count
        FROM words w
        JOIN words_groups wg ON w.id = wg.word_id
        WHERE wg.group_id = ?
        LIMIT ? OFFSET ?
    `, groupID, perPage, offset)
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
	err = DB.QueryRow(`
        SELECT COUNT(*) 
        FROM words w
        JOIN words_groups wg ON w.id = wg.word_id
        WHERE wg.group_id = ?
    `, groupID).Scan(&totalItems)
	if err != nil {
		return nil, 0, err
	}

	return words, totalItems, nil
}

func GetGroupStudySessions(groupID, page, perPage int) ([]map[string]interface{}, int, error) {
	offset := (page - 1) * perPage
	rows, err := DB.Query(`
        SELECT ss.id, 'Vocabulary Practice' AS activity_name, g.name AS group_name, ss.created_at AS start_time, ss.created_at AS end_time, 
               (SELECT COUNT(*) FROM word_review_items WHERE study_session_id = ss.id) AS number_of_review_items
        FROM study_sessions ss
        JOIN groups g ON ss.group_id = g.id
        WHERE ss.group_id = ?
        LIMIT ? OFFSET ?
    `, groupID, perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var sessions []map[string]interface{}
	for rows.Next() {
		var sessionID int
		var activityName, groupName string
		var startTime, endTime time.Time
		var numberOfReviewItems int
		if err := rows.Scan(&sessionID, &activityName, &groupName, &startTime, &endTime, &numberOfReviewItems); err != nil {
			return nil, 0, err
		}
		session := map[string]interface{}{
			"id":                     sessionID,
			"activity_name":          activityName,
			"group_name":             groupName,
			"start_time":             startTime,
			"end_time":               endTime,
			"number_of_review_items": numberOfReviewItems,
		}
		sessions = append(sessions, session)
	}

	var totalItems int
	err = DB.QueryRow(`
        SELECT COUNT(*) 
        FROM study_sessions 
        WHERE group_id = ?
    `, groupID).Scan(&totalItems)
	if err != nil {
		return nil, 0, err
	}

	return sessions, totalItems, nil
}

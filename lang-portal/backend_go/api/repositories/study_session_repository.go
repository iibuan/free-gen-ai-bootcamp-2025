package repositories

import (
	"backend_go/api/models"
	"time"
)

func GetStudySessions(page, perPage int) ([]map[string]interface{}, int, error) {
	offset := (page - 1) * perPage
	rows, err := DB.Query(`
		SELECT ss.id, 'Vocabulary Practice' AS activity_name, g.name AS group_name, ss.created_at AS start_time, ss.created_at AS end_time, 
			   (SELECT COUNT(*) FROM word_review_items WHERE study_session_id = ss.id) AS number_of_review_items
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		LIMIT ? OFFSET ?
	`, perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var sessions []map[string]interface{}
	for rows.Next() {
		var session map[string]interface{}
		var id int
		var activityName, groupName string
		var startTime, endTime time.Time
		var numberOfReviewItems int
		if err := rows.Scan(&id, &activityName, &groupName, &startTime, &endTime, &numberOfReviewItems); err != nil {
			return nil, 0, err
		}
		session = map[string]interface{}{
			"id":                     id,
			"activity_name":          activityName,
			"group_name":             groupName,
			"start_time":             startTime,
			"end_time":               endTime,
			"number_of_review_items": numberOfReviewItems,
		}
		sessions = append(sessions, session)
	}

	var totalItems int
	err = DB.QueryRow("SELECT COUNT(*) FROM study_sessions").Scan(&totalItems)
	if err != nil {
		return nil, 0, err
	}

	return sessions, totalItems, nil
}

func GetStudySession(id int) (map[string]interface{}, error) {
	var session map[string]interface{}
	var activityName, groupName string
	var startTime, endTime time.Time
	var numberOfReviewItems int

	err := DB.QueryRow(`
		SELECT 'Vocabulary Practice' AS activity_name, g.name AS group_name, ss.created_at AS start_time, ss.created_at AS end_time, 
			   (SELECT COUNT(*) FROM word_review_items WHERE study_session_id = ss.id) AS number_of_review_items
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		WHERE ss.id = ?
	`, id).Scan(&activityName, &groupName, &startTime, &endTime, &numberOfReviewItems)
	if err != nil {
		return nil, err
	}

	session = map[string]interface{}{
		"id":                     id,
		"activity_name":          activityName,
		"group_name":             groupName,
		"start_time":             startTime,
		"end_time":               endTime,
		"number_of_review_items": numberOfReviewItems,
	}

	return session, nil
}

func GetStudySessionWords(studySessionID, page, perPage int) ([]map[string]interface{}, int, error) {
	offset := (page - 1) * perPage
	rows, err := DB.Query(`
		SELECT w.id, w.bahasa_indonesia, w.english, wri.correct, wri.created_at
		FROM words w
		JOIN word_review_items wri ON w.id = wri.word_id
		WHERE wri.study_session_id = ?
		LIMIT ? OFFSET ?
	`, studySessionID, perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var words []map[string]interface{}
	for rows.Next() {
		var word models.Word
		var correct bool
		var createdAt time.Time
		if err := rows.Scan(&word.ID, &word.BahasaIndonesia, &word.English, &correct, &createdAt); err != nil {
			return nil, 0, err
		}
		wordMap := map[string]interface{}{
			"id":               word.ID,
			"bahasa_indonesia": word.BahasaIndonesia,
			"english":          word.English,
			"correct":          correct,
			"created_at":       createdAt,
		}
		words = append(words, wordMap)
	}

	var totalItems int
	err = DB.QueryRow(`
		SELECT COUNT(*) 
		FROM word_review_items 
		WHERE study_session_id = ?
	`, studySessionID).Scan(&totalItems)
	if err != nil {
		return nil, 0, err
	}

	return words, totalItems, nil
}

func CreateWordReviewItem(studySessionID, wordID int, correct bool) (map[string]interface{}, error) {
	createdAt := time.Now()
	res, err := DB.Exec(`
		INSERT INTO word_review_items (word_id, study_session_id, correct, created_at) 
		VALUES (?, ?, ?, ?)
	`, wordID, studySessionID, correct, createdAt)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	reviewItem := map[string]interface{}{
		"id":               id,
		"word_id":          wordID,
		"study_session_id": studySessionID,
		"correct":          correct,
		"created_at":       createdAt,
	}

	return reviewItem, nil
}

func GetLastStudySession() (map[string]interface{}, error) {
	var result map[string]interface{}

	rows, err := DB.Query(`
		SELECT ss.id, ss.created_at, 
			   (SELECT COUNT(*) FROM word_review_items WHERE study_session_id = ss.id AND correct = 1) AS correct_count,
			   (SELECT COUNT(*) FROM word_review_items WHERE study_session_id = ss.id AND correct = 0) AS wrong_count,
			   g.id, g.name
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		ORDER BY ss.created_at DESC LIMIT 1
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var id, correctCount, wrongCount, groupID int
		var groupName string
		var lastUsed time.Time

		if err := rows.Scan(&id, &lastUsed, &correctCount, &wrongCount, &groupID, &groupName); err != nil {
			return nil, err
		}

		result = map[string]interface{}{
			"id":            id,
			"activity_name": "Vocabulary Practice", // Assuming a static activity name for now
			"last_used":     lastUsed,
			"correct_count": correctCount,
			"wrong_count":   wrongCount,
			"group_id":      groupID,
			"group_name":    groupName,
		}
	}

	return result, nil
}

func GetStudyProgress() (map[string]interface{}, error) {
	var result map[string]interface{}

	// Query total words studied
	var totalWordsStudied int
	err := DB.QueryRow(`
		SELECT COUNT(DISTINCT word_id) FROM word_review_items
	`).Scan(&totalWordsStudied)
	if err != nil {
		return nil, err
	}

	// Query mastery progress
	var masteryProgress float64
	err = DB.QueryRow(`
		SELECT AVG(correct) FROM word_review_items
	`).Scan(&masteryProgress)
	if err != nil {
		return nil, err
	}

	result = map[string]interface{}{
		"total_words_studied": totalWordsStudied,
		"mastery_progress":    masteryProgress,
	}

	return result, nil
}

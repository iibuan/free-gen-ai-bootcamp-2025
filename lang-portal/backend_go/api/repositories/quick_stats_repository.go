package repositories

func GetQuickStats() (map[string]interface{}, error) {
	var result map[string]interface{}

	// Query success rate
	var successRate int
	err := DB.QueryRow(`
        SELECT (SUM(correct) * 100 / COUNT(*)) FROM word_review_items
    `).Scan(&successRate)
	if err != nil {
		return nil, err
	}

	// Query total study sessions
	var totalStudySessions int
	err = DB.QueryRow(`
        SELECT COUNT(*) FROM study_sessions
    `).Scan(&totalStudySessions)
	if err != nil {
		return nil, err
	}

	// Query total active groups
	var totalActiveGroups int
	err = DB.QueryRow(`
        SELECT COUNT(DISTINCT group_id) FROM study_sessions
    `).Scan(&totalActiveGroups)
	if err != nil {
		return nil, err
	}

	// Query study streak
	var studyStreak int
	err = DB.QueryRow(`
        SELECT COUNT(DISTINCT DATE(created_at)) FROM study_sessions
    `).Scan(&studyStreak)
	if err != nil {
		return nil, err
	}

	result = map[string]interface{}{
		"success_rate":         successRate,
		"total_study_sessions": totalStudySessions,
		"total_active_groups":  totalActiveGroups,
		"study_streak":         studyStreak,
	}

	return result, nil
}

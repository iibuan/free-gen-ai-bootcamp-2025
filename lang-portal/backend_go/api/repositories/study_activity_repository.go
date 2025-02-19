package repositories

import (
	"backend_go/api/models"
	"time"
)

func GetStudyActivities() ([]models.StudyActivity, error) {
	rows, err := DB.Query("SELECT id, study_session_id, group_id, created_at FROM study_activities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []models.StudyActivity
	for rows.Next() {
		var activity models.StudyActivity
		if err := rows.Scan(&activity.ID, &activity.StudySessionID, &activity.GroupID, &activity.CreatedAt); err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}

func GetStudyActivity(id string) (models.StudyActivity, error) {
	var activity models.StudyActivity
	err := DB.QueryRow("SELECT id, study_session_id, group_id, created_at FROM study_activities WHERE id = ?", id).
		Scan(&activity.ID, &activity.StudySessionID, &activity.GroupID, &activity.CreatedAt)
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func CreateStudyActivity(studySessionID, groupID int) (models.StudyActivity, error) {
	var activity models.StudyActivity
	createdAt := time.Now()
	res, err := DB.Exec("INSERT INTO study_activities (study_session_id, group_id, created_at) VALUES (?, ?, ?)", studySessionID, groupID, createdAt)
	if err != nil {
		return activity, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return activity, err
	}

	activity = models.StudyActivity{
		ID:             int(id),
		StudySessionID: studySessionID,
		GroupID:        groupID,
		CreatedAt:      createdAt,
	}
	return activity, nil
}

package services

import (
	"backend_go/api/models"
	"backend_go/api/repositories"
)

type StudyActivityService struct{}

func (s *StudyActivityService) GetStudyActivities() ([]models.StudyActivity, error) {
	return repositories.GetStudyActivities()
}

func (s *StudyActivityService) GetStudyActivity(id string) (models.StudyActivity, error) {
	return repositories.GetStudyActivity(id)
}

func (s *StudyActivityService) CreateStudyActivity(studySessionID, groupID int) (models.StudyActivity, error) {
	return repositories.CreateStudyActivity(studySessionID, groupID)
}

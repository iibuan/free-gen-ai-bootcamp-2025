package services

import (
	"backend_go/api/repositories"
)

type StudySessionService struct{}

func (s *StudySessionService) GetStudySessions(page, perPage int) ([]map[string]interface{}, int, error) {
	return repositories.GetStudySessions(page, perPage)
}

func (s *StudySessionService) GetStudySession(id int) (map[string]interface{}, error) {
	return repositories.GetStudySession(id)
}

func (s *StudySessionService) GetStudySessionWords(studySessionID, page, perPage int) ([]map[string]interface{}, int, error) {
	return repositories.GetStudySessionWords(studySessionID, page, perPage)
}

func (s *StudySessionService) CreateWordReviewItem(studySessionID, wordID int, correct bool) (map[string]interface{}, error) {
	return repositories.CreateWordReviewItem(studySessionID, wordID, correct)
}

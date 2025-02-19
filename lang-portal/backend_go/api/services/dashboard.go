package services

import (
	"backend_go/api/repositories"
)

type DashboardService struct{}

func (s *DashboardService) GetLastStudySession() (map[string]interface{}, error) {
	return repositories.GetLastStudySession()
}

func (s *DashboardService) GetStudyProgress() (map[string]interface{}, error) {
	return repositories.GetStudyProgress()
}

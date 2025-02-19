package services

import (
	"backend_go/api/repositories"
)

type QuickStatsService struct{}

func (s *QuickStatsService) GetQuickStats() (map[string]interface{}, error) {
	return repositories.GetQuickStats()
}

package services

import (
	"backend_go/api/repositories"
)

type ResetService struct{}

func (s *ResetService) ResetHistory() error {
	return repositories.ResetHistory()
}

func (s *ResetService) FullReset() error {
	return repositories.FullReset()
}

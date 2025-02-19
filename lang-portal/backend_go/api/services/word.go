package services

import (
	"backend_go/api/repositories"
)

type WordService struct{}

func (s *WordService) GetWords(page, perPage int) ([]map[string]interface{}, int, error) {
	return repositories.GetWords(page, perPage)
}

func (s *WordService) GetWord(id int) (map[string]interface{}, error) {
	return repositories.GetWord(id)
}

package services

import (
	"backend_go/api/repositories"
)

type GroupService struct{}

func (s *GroupService) GetGroups(page, perPage int) ([]map[string]interface{}, int, error) {
	return repositories.GetGroups(page, perPage)
}

func (s *GroupService) GetGroup(id int) (map[string]interface{}, error) {
	return repositories.GetGroup(id)
}

func (s *GroupService) GetGroupWords(groupID, page, perPage int) ([]map[string]interface{}, int, error) {
	return repositories.GetGroupWords(groupID, page, perPage)
}

func (s *GroupService) GetGroupStudySessions(groupID, page, perPage int) ([]map[string]interface{}, int, error) {
	return repositories.GetGroupStudySessions(groupID, page, perPage)
}

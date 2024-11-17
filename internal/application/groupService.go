package application

import "task-plan/internal/application/requestModels"

type GroupService struct {
	repo *ITaskService
}

func NewGroupService(repo *ITaskService) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) Create(add requestModels.GroupToAdd) (int, error) {
	return 0, nil
}

func (s *GroupService) GetById(id int) (requestModels.GroupToGet, error) {
	return requestModels.GroupToGet{}, nil
}

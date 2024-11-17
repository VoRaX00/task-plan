package application

import (
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
	"task-plan/pkg/mapper"
)

type GroupService struct {
	repo   IGroupRepository
	mapper *mapper.GroupMapper
}

func NewGroupService(repo IGroupRepository) *GroupService {
	return &GroupService{
		repo:   repo,
		mapper: mapper.NewGroupMapper(),
	}
}

func (s *GroupService) Create(groupAdd requestModels.GroupToAdd) (int, error) {
	group := s.mapper.GroupAddToGroup(groupAdd)
	return s.repo.Create(group)
}

func (s *GroupService) GetById(id int) (domain.Group, error) {
	return s.repo.GetById(id)
}

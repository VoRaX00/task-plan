package mapper

import (
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
)

type GroupMapper struct {
}

func NewGroupMapper() *GroupMapper {
	return &GroupMapper{}
}

func (m *GroupMapper) GroupAddToGroup(group requestModels.GroupToAdd) domain.Group {
	return domain.Group{
		Name:          group.Name,
		Users:         []domain.User{group.AdminUser},
		LevelProgress: group.LevelProgress,
	}
}

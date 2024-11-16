package infrastructure

import (
	"gorm.io/gorm"
	"task-plan/internal/domain"
)

type GroupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) *GroupRepo {
	return &GroupRepo{
		db: db,
	}
}

func (r *GroupRepo) Create(group domain.Group) (int, error) {
	tx := r.db.Create(group)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return 1, nil
}

func (r *GroupRepo) GetById(id int) (domain.Group, error) {
	var group domain.Group
	if tx := r.db.First(&group, id); tx.Error != nil {
		return domain.Group{}, tx.Error
	}
	return group, nil
}

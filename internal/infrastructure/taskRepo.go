package infrastructure

import (
	"gorm.io/gorm"
	"task-plan/internal/domain"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (r *TaskRepo) Create(task domain.Task) (int, error) {
	tx := r.db.Create(&task)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return 1, nil
}

func (r *TaskRepo) GetById(id int) (domain.Task, error) {
	var task domain.Task
	tx := r.db.First(&task, id)
	if tx.Error != nil {
		return task, tx.Error
	}
	return task, nil
}

func (r *TaskRepo) GetByGroupId(id int) ([]domain.Task, error) {
	var task []domain.Task
	tx := r.db.Where("group_id = ?", id).Find(&task)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return task, nil
}

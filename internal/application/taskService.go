package application

import (
	"task-plan/internal/domain"
)

type TaskService struct {
	repo ITaskRepository
}

func NewTaskService(repo ITaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(task domain.Task) (int, error) {
	return s.repo.Create(task)
}

func (s *TaskService) GetById(id int) (domain.Task, error) {
	return s.repo.GetById(id)
}

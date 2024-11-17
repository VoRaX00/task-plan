package application

import "task-plan/internal/application/requestModels"

type TaskService struct {
	repo ITaskRepository
}

func NewTaskService(repo ITaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(task requestModels.TaskToAdd) (int, error) {
	return 0, nil
}

func (s *TaskService) GetById(id int) (requestModels.TaskToGet, error) {
	return requestModels.TaskToGet{}, nil
}

package application

type TaskService struct {
	repo *ITaskRepository
}

func NewTaskService(repo *ITaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

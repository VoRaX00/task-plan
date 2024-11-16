package application

type GroupService struct {
	repo *ITaskService
}

func NewGroupService(repo *ITaskService) *GroupService {
	return &GroupService{repo: repo}
}

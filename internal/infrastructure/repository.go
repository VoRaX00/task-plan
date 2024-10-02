package infrastructure

type Repository struct {
	IAuthRepository
}

func NewRepository() *Repository {
	return &Repository{
		IAuthRepository: NewAuthRepository(),
	}
}

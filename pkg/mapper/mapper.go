package mapper

type Mapper struct {
	IUserMapper
}

func NewMapper() *Mapper {
	return &Mapper{
		IUserMapper: NewUserMapper(),
	}
}

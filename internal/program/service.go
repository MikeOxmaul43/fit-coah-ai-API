package program

type Service struct {
	*Repository
}

func NewProgramService(repository *Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

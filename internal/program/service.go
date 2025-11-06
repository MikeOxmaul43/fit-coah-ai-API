package program

import "sportTrackerAPI/internal/user"

type Service struct {
	*Repository
	UserRepository *user.Repository
}

func NewProgramService(repository *Repository, userRepository *user.Repository) *Service {
	return &Service{
		Repository:     repository,
		UserRepository: userRepository,
	}
}

func (service *Service) Create(request CreateProgramRequest, userEmail string) error {
	creator, err := service.UserRepository.FindByEmail(userEmail)
	if err != nil {
		return err
	}

	program := Program{
		Title:         request.Title,
		Description:   request.Description,
		Level:         request.Level,
		DurationWeeks: request.DurationWeeks,
		CreatorType:   "user",
		CreatedBy:     &creator.ID,
	}
	_, err = service.Repository.Create(program)
	return err
}

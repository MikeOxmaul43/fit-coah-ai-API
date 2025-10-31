package exercise

import (
	pkg "sportTrackerAPI/db"
)

type Repository struct {
	DataBase *pkg.Db
}

func NewExerciseRepository(dataBase *pkg.Db) *Repository { return &Repository{DataBase: dataBase} }

func (repo *Repository) FindByName(name string) (*Exercise, error) {
	var exercise Exercise
	result := repo.DataBase.First(&exercise, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &exercise, nil
}

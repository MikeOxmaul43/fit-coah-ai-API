package exercise

import (
	pkg "sportTrackerAPI/db"
)

type Repository struct {
	DataBase *pkg.Db
}

func NewExerciseRepository(dataBase *pkg.Db) *Repository { return &Repository{DataBase: dataBase} }

func (repo *Repository) GetAll() ([]Exercise, error) {
	var exercises []Exercise
	err := repo.DataBase.DB.Find(&exercises).Error
	if err != nil {
		return nil, err
	}
	return exercises, nil
}

func (repo Repository) Update(exercise Exercise) error {
	return repo.DataBase.DB.Save(exercise).Error
}

func (repo Repository) Delete(id uint) error {
	return repo.DataBase.DB.Delete(&Exercise{}, id).Error
}
func (repo *Repository) GetByName(name string) (*Exercise, error) {
	var exercise Exercise
	result := repo.DataBase.First(&exercise, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &exercise, nil
}

func (repo *Repository) GetByMuscleGroup(muscleGroup string) ([]Exercise, error) {
	var exercises []Exercise
	result := repo.DataBase.Find(&exercises, "muscle_group = ?", muscleGroup)
	if result.Error != nil {
		return nil, result.Error
	}
	return exercises, nil
}

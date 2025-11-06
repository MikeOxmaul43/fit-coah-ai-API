package program

import (
	"gorm.io/gorm/clause"
	"sportTrackerAPI/db"
)

type Repository struct {
	DataBase *db.Db
}

func NewProgramRepository(dataBase *db.Db) *Repository { return &Repository{DataBase: dataBase} }

func (repo *Repository) Create(program Program) (*Program, error) {
	result := repo.DataBase.DB.Create(&program)
	if result.Error != nil {
		return nil, result.Error
	}
	return &program, nil
}

func (repo *Repository) Delete(programID uint) error {
	return repo.DataBase.DB.Delete(&Program{}, programID).Error
}

func (repo *Repository) Update(program Program) error {
	return repo.DataBase.Clauses(clause.Returning{}).Updates(program).Error
}

func (repo *Repository) GetById(id uint) (*Program, error) {
	var program Program
	result := repo.DataBase.DB.First(&program, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &program, nil
}

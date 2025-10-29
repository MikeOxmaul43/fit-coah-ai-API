package user

import (
	pkg "sportTrackerAPI/db"
)

type Repository struct {
	DataBase *pkg.Db
}

func NewUserRepository(dataBase *pkg.Db) *Repository { return &Repository{DataBase: dataBase} }

func (repo *Repository) Create(user *User) (*User, error) {
	result := repo.DataBase.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *Repository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.DataBase.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

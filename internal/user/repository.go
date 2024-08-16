package user

import "gorm.io/gorm"

type Repository interface {
	GetUserByID(id uint) (User, error)
}

type repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return &repository{database: database}
}

func (r *repository) GetUserByID(id uint) (User, error) {
	usr := User{}
	err := r.database.First(&usr, "id = ?", id).Error

	if err != nil {
		return User{}, err
	}

	return usr, nil	
}
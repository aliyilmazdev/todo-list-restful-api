package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByID(id string) (User, error)
	GetUserByEmail(email string) (User, error)
	Create(user *User) error 
 }

type repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return &repository{database: database}
}

func (r *repository) Create(user *User) error {
	err := r.database.Create(user).Error

	if err != nil {
		return err
	}

	return nil
} 

func (r *repository) GetUserByID(id string) (User, error) {
	usr := User{}
	err := r.database.First(&usr, "id = ?", id).Error

	if err != nil {
		return User{}, err
	}

	return usr, nil	
}

func (r *repository) GetUserByEmail(email string) (User, error) {
	usr := User{}
	err := r.database.First(&usr, "email = ?", email).Error

	if err != nil {
		return User{}, err
	}

	return usr, nil
}
package user

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name	 string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique;not null"`
}

type RegisterRequest struct {
	Email 	 string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name 	 string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
}

type LoginRequest struct {
	Email 	 string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
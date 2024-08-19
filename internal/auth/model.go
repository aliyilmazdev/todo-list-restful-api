package auth

type RegisterRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name 	 string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
}

type LoginRequest struct {
	Email 	 string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
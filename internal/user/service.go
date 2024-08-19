package user

type Service interface {
	GetUserByID(id string) (User, error)
	GetUserByEmail(email string) (User, error)
	Create(request RegisterRequest) (*User ,error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(request RegisterRequest) (*User, error) {
	
	u := User{
		Email: request.Email,
		Password: request.Password,
		Name: request.Name,
		Surname: request.Surname,
	}

	err := s.repository.Create(&u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *service) GetUserByID(id string) (User, error) {
	u, err := s.repository.GetUserByID(id)

	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *service) GetUserByEmail(email string) (User, error) {
	u, err := s.repository.GetUserByEmail(email)

	if err != nil {
		return User{}, err
	}

	return u, nil
}
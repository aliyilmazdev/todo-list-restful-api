package user

type Service interface {
	GetUserByID(id uint) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetUserByID(id uint) (User, error) {
	u, err := s.repository.GetUserByID(id)

	if err != nil {
		return User{}, err
	}

	return u, nil
}
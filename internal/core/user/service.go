package user

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateUser(u User) (User, error) {
	return s.repo.Create(u)
}

func (s *Service) GetUser(id int) (User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) UpdateUser(u User) (User, error) {
	return s.repo.Update(u)
}
func (s *Service) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
func (s *Service) ListUsers() ([]User, error) {
	return s.repo.FindAll()
}

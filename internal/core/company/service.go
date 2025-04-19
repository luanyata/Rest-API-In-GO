package company

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}
func (s *Service) CreateCompany(c Company) (Company, error) {
	return s.repo.Create(c)
}
func (s *Service) GetCompany(id int) (Company, error) {
	return s.repo.FindByID(id)
}
func (s *Service) UpdateCompany(c Company) (Company, error) {
	return s.repo.Update(c)
}
func (s *Service) DeleteCompany(id int) error {
	return s.repo.Delete(id)
}
func (s *Service) ListCompanies() ([]Company, error) {
	return s.repo.FindAll()
}

package user

import (
	"net/http"
	"rest-go/internal/infra/security"
	"rest-go/internal/shared/errs"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateUser(u User) (User, error) {
	if !u.IsValidEmail() {
		return User{}, errs.BadRequestException("Email inválido", "invalid_email")
	}

	u.Normalize()

	_, err := s.GetUserByEmail(u.Email)

	if err == nil {
		return User{}, errs.ConflictException("Email já cadastrado", "email_already_exists")
	}

	if httpErr, ok := err.(errs.HttpErrorInterface); !ok || httpErr.StatusCode() != http.StatusNotFound {
		return User{}, err
	}

	hashedPassword, err := security.HashPassword(u.Password)
	if err != nil {
		return User{}, errs.InternalServerException("Erro ao criar usuário", "hash_error")
	}

	u.Password = hashedPassword

	return s.repo.Create(u)
}

func (s *Service) GetUser(id int) (User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetUserByEmail(email string) (User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return User{}, err
	}

	return u, nil
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

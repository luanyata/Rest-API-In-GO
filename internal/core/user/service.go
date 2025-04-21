package user

import (
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
		return User{}, errs.BadRequestException(
			errs.UserError.EmailInvalid.Message,
			errs.UserError.EmailInvalid.Type,
		)
	}

	u.Normalize()

	_, found := s.repo.FindByEmail(u.Email)

	if found {
		return User{}, errs.ConflictException(
			errs.UserError.EmailAlreadyExists.Message,
			errs.UserError.EmailAlreadyExists.Type,
		)
	}

	hashedPassword, err := security.HashPassword(u.Password)
	if err != nil {
		return User{}, errs.InternalServerException(
			errs.UserError.FailHashPassword.Message,
			errs.UserError.FailHashPassword.Type,
		)
	}

	u.Password = hashedPassword

	return s.repo.Create(u)
}

func (s *Service) GetUser(id int) (User, error) {

	u, found := s.repo.FindByID(id)

	if !found {
		return User{}, errs.NotFoundException(
			errs.UserError.NotFound.Message,
			errs.UserError.NotFound.Type,
		)
	}
	return u, nil
}

func (s *Service) GetUserByEmail(email string) (User, bool) {
	u, found := s.repo.FindByEmail(email)

	return u, found
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

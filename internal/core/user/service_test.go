package user

import (
	"rest-go/internal/shared/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	repo := NewMemoryRepo()
	service := NewService(repo)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		Password:  "password123",
	}
	createdUser, err := service.CreateUser(user)
	assert.NoError(t, err, "expected no error but got one")
	assert.Equal(t, user.FirstName, createdUser.FirstName)
	assert.Equal(t, user.LastName, createdUser.LastName)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.NotEmpty(t, createdUser.ID, "expected user ID to be set")
	assert.NotEqual(t, user.Password, createdUser.Password, "expected password to be hashed")
	assert.NotEmpty(t, createdUser.Password, "expected hashed password to be set")
}

func TestCreateUser_EmailInvalid(t *testing.T) {
	repo := NewMemoryRepo()
	service := NewService(repo)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "invalid-email",
		Password:  "password123",
	}

	_, err := service.CreateUser(user)

	// Verificando se ocorreu um erro
	assert.Error(t, err, "expected error but got nil")

	// Verificando se o erro é do tipo esperado
	httpErr, ok := err.(errs.HttpErrorInterface)
	assert.True(t, ok, "expected HttpErrorInterface")

	// Verificando as propriedades do erro
	assert.Equal(t, 400, httpErr.StatusCode()) // status 400 (Bad Request)
	assert.Equal(t, errs.UserError.EmailInvalid.Type, httpErr.Type())
	assert.Equal(t, errs.UserError.EmailInvalid.Message, httpErr.Error())
}

func TestCreateUser_EmailAlreadyExists(t *testing.T) {
	repo := NewMemoryRepo()

	// Pré-populando o repositório com um usuário
	repo.Create(User{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "john@doe.com",
		Password:  "hashed-password",
	})

	service := NewService(repo)

	// Tentando criar um novo usuário com o mesmo email
	newUser := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		Password:  "password123",
	}
	_, err := service.CreateUser(newUser)

	// Verificando se ocorreu um erro
	assert.Error(t, err, "expected error but got nil")

	// Verificando se o erro é do tipo esperado
	httpErr, ok := err.(errs.HttpErrorInterface)
	assert.True(t, ok, "expected HttpErrorInterface")

	// Verificando as propriedades do erro
	assert.Equal(t, 409, httpErr.StatusCode()) // status 409 (Conflict)
	assert.Equal(t, errs.UserError.EmailAlreadyExists.Type, httpErr.Type())
	assert.Equal(t, errs.UserError.EmailAlreadyExists.Message, httpErr.Error())
}

func TestGetUser_Success(t *testing.T) {
	repo := NewMemoryRepo()
	service := NewService(repo)

	// Pré-populando o repositório com um usuário
	expectedUser := User{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "john@doe.com",
		Password:  "hashed-password",
	}
	repo.Create(expectedUser)
	// Buscando o usuário pelo ID
	user, err := service.GetUser(expectedUser.ID)
	assert.NoError(t, err, "expected no error but got one")
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.FirstName, user.FirstName)
	assert.Equal(t, expectedUser.LastName, user.LastName)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Password, user.Password)
}

func TestGetUser_NotFound(t *testing.T) {
	repo := NewMemoryRepo()
	service := NewService(repo)

	// Buscando um usuário que não existe
	user, err := service.GetUser(999)
	assert.Error(t, err, "expected error but got nil")
	assert.Equal(t, User{}, user, "expected empty user but got one")
	// Verificando se o erro é do tipo esperado
	httpErr, ok := err.(errs.HttpErrorInterface)
	assert.True(t, ok, "expected HttpErrorInterface")
	// Verificando as propriedades do erro
	assert.Equal(t, 404, httpErr.StatusCode()) // status 404 (Not Found)
	assert.Equal(t, errs.UserError.NotFound.Type, httpErr.Type())
	assert.Equal(t, errs.UserError.NotFound.Message, httpErr.Error())
}

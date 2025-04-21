package user

type Repository interface {
	Create(User) (User, error)
	FindAll() ([]User, error)
	FindByID(id int) (User, bool)
	FindByEmail(email string) (User, bool)
	Update(user User) (User, error)
	Delete(id int) error
}

package user

type Repository interface {
	Create(User) (User, error)
	FindAll() ([]User, error)
	FindByID(id int) (User, error)
	Update(user User) (User, error)
	Delete(id int) error
}

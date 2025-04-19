package company

type Repository interface {
	Create(Company) (Company, error)
	FindAll() ([]Company, error)
	FindByID(id int) (Company, error)
	Update(company Company) (Company, error)
	Delete(id int) error
}

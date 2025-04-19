package company

type memoryRepo struct {
	data   map[int]Company
	nextID int
}

func NewMemoryRepo() Repository {
	return &memoryRepo{
		data:   make(map[int]Company),
		nextID: 1,
	}
}

func (m *memoryRepo) FindAll() ([]Company, error) {
	companies := make([]Company, 0, len(m.data))
	for _, company := range m.data {
		companies = append(companies, company)
	}
	return companies, nil
}

func (m *memoryRepo) Create(c Company) (Company, error) {
	c.ID = m.nextID
	m.nextID++
	m.data[c.ID] = c
	return c, nil
}

func (m *memoryRepo) FindByID(id int) (Company, error) {
	c := m.data[id]
	return c, nil
}

func (m *memoryRepo) Update(company Company) (Company, error) {
	m.data[company.ID] = company
	return company, nil
}

func (m *memoryRepo) Delete(id int) error {
	delete(m.data, id)
	return nil
}

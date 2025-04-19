package user

type memoryRepo struct {
	data   map[int]User
	nextID int
}

func NewMemoryRepo() Repository {
	return &memoryRepo{
		data:   make(map[int]User),
		nextID: 1,
	}
}

func (m *memoryRepo) FindAll() ([]User, error) {
	users := make([]User, 0, len(m.data))
	for _, user := range m.data {
		users = append(users, user)
	}
	return users, nil
}

func (m *memoryRepo) Create(u User) (User, error) {
	u.ID = m.nextID
	m.nextID++
	m.data[u.ID] = u
	return u, nil
}

func (m *memoryRepo) FindByID(id int) (User, error) {
	u := m.data[id]
	return u, nil
}

func (m *memoryRepo) Update(user User) (User, error) {
	m.data[user.ID] = user
	return user, nil
}

func (m *memoryRepo) Delete(id int) error {
	delete(m.data, id)
	return nil
}

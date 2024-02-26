package chapter7

/* Package Provider */

type Store struct {
	db string
}

func NewDB() *Store {
	return &Store{
		db: "",
	}
}

func (s *Store) Insert(item any) error {
	return nil
}

func (s *Store) Get(id int) error {
	return nil
}

/* Package Comsumer */

type UserStorer interface {
	Insert(item any) error
	Get(id int) error
}

type UserService struct {
	store UserStorer
}

/* Accept inferfaces, return structs */
func NewUS(s UserStorer) *UserService {
	return &UserService{
		store: s,
	}
}

func (u *UserService) Dispatch()  {
	//....
}

func (u *UserService) CreateUser() {
	u.store.Insert(struct{ s string }{s: ""})
}

func (u *UserService) RetriveUser() {
	u.store.Get(10)
}



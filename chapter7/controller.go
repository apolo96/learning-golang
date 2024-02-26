package chapter7

type AsynStore interface {
	Dispatch() 
}

func boot() {
	var db = NewDB()
	var uss *UserService = NewUS(db)
	use := &UserService{
		store: db,
	}
	use.RetriveUser()
	uss.CreateUser()	
}

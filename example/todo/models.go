package todo

type Ownable interface {
	Owner() *User
}

type Todo struct {
	ID    int
	Text  string
	Done  bool
	owner *User
}

func (t Todo) Owner() *User {
	return t.owner
}

type User struct {
	ID   int
	Name string
}

type TodoNew struct {
	ID   int
	Text []string
}

type TodoNew1 struct {
	ID   int
	Text []string
}

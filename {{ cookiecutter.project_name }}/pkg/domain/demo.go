package domain

type DemoRepository interface {
	Store(Demo) error
	LoadByID(string) (*Demo, error)
}

type Demo struct {
	ID    string
	Label string
}

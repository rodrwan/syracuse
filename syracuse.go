package syracuse

// Citizen ...
type Citizen struct {
	ID       string
	Email    string
	Fullname string
}

// Citizens ...
type Citizens interface {
	Get(string) (*Citizen, error)
	Select() ([]*Citizen, error)
	Create(*Citizen) error
	Update(*Citizen) error
	Delete(*Citizen) error
}

package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/rodrwan/syracuse"
)

// CitizensService ...
type CitizensService struct {
	Store *sqlx.DB
}

// Get ...
func (cs *CitizensService) Get(ID string) (*syracuse.Citizen, error) {
	sql := "Select * from users where id = ? and deleted_at is null"

	rows, err := cs.Store.Queryx(sql)
	if err != nil {
		return nil, err
	}

	c := &syracuse.Citizen{}
	if err := rows.Scan(c); err != nil {
		return nil, err
	}

	return c, nil
}

// Select ...
func (cs *CitizensService) Select() ([]*syracuse.Citizen, error) {
	return nil, nil
}

// Create ...
func (cs *CitizensService) Create(c *syracuse.Citizen) error {
	return nil
}

// Update ...
func (cs *CitizensService) Update(*syracuse.Citizen) error {
	return nil
}

// Delete ...
func (cs *CitizensService) Delete(*syracuse.Citizen) error {
	return nil
}

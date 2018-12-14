package postgres

import (
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rodrwan/syracuse"
)

// CitizensService ...
type CitizensService struct {
	Store *sqlx.DB
}

// Get ...
func (cs *CitizensService) Get(ID string) (*syracuse.Citizen, error) {
	query := squirrel.Select("*").From("users").Where("id = ?", ID).Where("deleted_at is null")

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := cs.Store.QueryRowx(sql, args...)

	c := &syracuse.Citizen{}
	if err := row.StructScan(c); err != nil {
		return nil, err
	}

	return c, nil
}

// Select ...
func (cs *CitizensService) Select() ([]*syracuse.Citizen, error) {
	query := squirrel.Select("*").From("users")

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := cs.Store.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	cc := make([]*syracuse.Citizen, 0)

	for rows.Next() {
		c := &syracuse.Citizen{}
		if err := rows.StructScan(c); err != nil {
			return nil, err
		}
		cc = append(cc, c)
	}

	return cc, nil
}

// Create ...
func (cs *CitizensService) Create(c *syracuse.Citizen) error {
	sql, args, err := squirrel.
		Insert("users").
		Columns("email", "fullname").
		Values(c.Email, c.Fullname).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := cs.Store.QueryRowx(sql, args...)
	if err := row.StructScan(c); err != nil {
		return err
	}

	return nil
}

// Update ...
func (cs *CitizensService) Update(c *syracuse.Citizen) error {
	sql, args, err := squirrel.Update("users").
		Set("email", c.Email).
		Set("fullname", c.Fullname).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := cs.Store.QueryRowx(sql, args...)
	return row.StructScan(c)
}

// Delete ...
func (cs *CitizensService) Delete(c *syracuse.Citizen) error {
	row := cs.Store.QueryRowx(
		"update users set deleted_at = $1 where id = $2 returning *",
		time.Now(), c.ID,
	)

	if err := row.StructScan(c); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}

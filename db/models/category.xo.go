// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

// Category represents a row from 'public.categories'.
type Category struct {
	Snowflake   int64          `json:"snowflake"`   // snowflake
	CreatedAt   *time.Time     `json:"created_at"`  // created_at
	DeletedAt   pq.NullTime    `json:"deleted_at"`  // deleted_at
	Title       string         `json:"title"`       // title
	Description sql.NullString `json:"description"` // description
	Color       sql.NullInt64  `json:"color"`       // color

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Category exists in the database.
func (c *Category) Exists() bool {
	return c._exists
}

// Deleted provides information if the Category has been deleted from the database.
func (c *Category) Deleted() bool {
	return c._deleted
}

// Insert inserts the Category to the database.
func (c *Category) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.categories (` +
		`snowflake, created_at, deleted_at, title, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`

	// run query
	XOLog(sqlstr, c.Snowflake, c.CreatedAt, c.DeletedAt, c.Title, c.Description, c.Color)
	err = db.QueryRow(sqlstr, c.Snowflake, c.CreatedAt, c.DeletedAt, c.Title, c.Description, c.Color).Scan(&c.Snowflake)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Category in the database.
func (c *Category) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.categories SET (` +
		`created_at, deleted_at, title, description, color` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE snowflake = $6`

	// run query
	XOLog(sqlstr, c.CreatedAt, c.DeletedAt, c.Title, c.Description, c.Color, c.Snowflake)
	_, err = db.Exec(sqlstr, c.CreatedAt, c.DeletedAt, c.Title, c.Description, c.Color, c.Snowflake)
	return err
}

// Save saves the Category to the database.
func (c *Category) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Category.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Category) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.categories (` +
		`snowflake, created_at, deleted_at, title, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (snowflake) DO UPDATE SET (` +
		`snowflake, created_at, deleted_at, title, description, color` +
		`) = (` +
		`EXCLUDED.snowflake, EXCLUDED.created_at, EXCLUDED.deleted_at, EXCLUDED.title, EXCLUDED.description, EXCLUDED.color` +
		`)`

	// run query
	XOLog(sqlstr, c.Snowflake, c.CreatedAt, c.DeletedAt, c.Title, c.Description, c.Color)
	_, err = db.Exec(sqlstr, c.Snowflake, c.CreatedAt, c.DeletedAt, c.Title, c.Description, c.Color)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Category from the database.
func (c *Category) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.categories WHERE snowflake = $1`

	// run query
	XOLog(sqlstr, c.Snowflake)
	_, err = db.Exec(sqlstr, c.Snowflake)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// CategoryBySnowflake retrieves a row from 'public.categories' as a Category.
//
// Generated from index 'categories_pkey'.
func CategoryBySnowflake(db XODB, snowflake int64) (*Category, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`snowflake, created_at, deleted_at, title, description, color ` +
		`FROM public.categories ` +
		`WHERE snowflake = $1`

	// run query
	XOLog(sqlstr, snowflake)
	c := Category{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, snowflake).Scan(&c.Snowflake, &c.CreatedAt, &c.DeletedAt, &c.Title, &c.Description, &c.Color)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// CategoriesByTitle retrieves a row from 'public.categories' as a Category.
//
// Generated from index 'categories_title_index'.
func CategoriesByTitle(db XODB, title string) ([]*Category, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`snowflake, created_at, deleted_at, title, description, color ` +
		`FROM public.categories ` +
		`WHERE title = $1`

	// run query
	XOLog(sqlstr, title)
	q, err := db.Query(sqlstr, title)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Category{}
	for q.Next() {
		c := Category{
			_exists: true,
		}

		// scan
		err = q.Scan(&c.Snowflake, &c.CreatedAt, &c.DeletedAt, &c.Title, &c.Description, &c.Color)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}

// CategoryByTitle retrieves a row from 'public.categories' as a Category.
//
// Generated from index 'categories_title_key'.
func CategoryByTitle(db XODB, title string) (*Category, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`snowflake, created_at, deleted_at, title, description, color ` +
		`FROM public.categories ` +
		`WHERE title = $1`

	// run query
	XOLog(sqlstr, title)
	c := Category{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, title).Scan(&c.Snowflake, &c.CreatedAt, &c.DeletedAt, &c.Title, &c.Description, &c.Color)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
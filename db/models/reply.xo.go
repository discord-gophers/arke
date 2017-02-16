// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

// Reply represents a row from 'public.replies'.
type Reply struct {
	Snowflake int64         `json:"snowflake"`  // snowflake
	CreatedAt *time.Time    `json:"created_at"` // created_at
	DeletedAt pq.NullTime   `json:"deleted_at"` // deleted_at
	AuthorID  sql.NullInt64 `json:"author_id"`  // author_id
	Body      string        `json:"body"`       // body
	ParentID  sql.NullInt64 `json:"parent_id"`  // parent_id
	TopicID   int64         `json:"topic_id"`   // topic_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Reply exists in the database.
func (r *Reply) Exists() bool {
	return r._exists
}

// Deleted provides information if the Reply has been deleted from the database.
func (r *Reply) Deleted() bool {
	return r._deleted
}

// Insert inserts the Reply to the database.
func (r *Reply) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.replies (` +
		`snowflake, created_at, deleted_at, author_id, body, parent_id, topic_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)`

	// run query
	XOLog(sqlstr, r.Snowflake, r.CreatedAt, r.DeletedAt, r.AuthorID, r.Body, r.ParentID, r.TopicID)
	err = db.QueryRow(sqlstr, r.Snowflake, r.CreatedAt, r.DeletedAt, r.AuthorID, r.Body, r.ParentID, r.TopicID).Scan(&r.Snowflake)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Update updates the Reply in the database.
func (r *Reply) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.replies SET (` +
		`created_at, deleted_at, author_id, body, parent_id, topic_id` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE snowflake = $7`

	// run query
	XOLog(sqlstr, r.CreatedAt, r.DeletedAt, r.AuthorID, r.Body, r.ParentID, r.TopicID, r.Snowflake)
	_, err = db.Exec(sqlstr, r.CreatedAt, r.DeletedAt, r.AuthorID, r.Body, r.ParentID, r.TopicID, r.Snowflake)
	return err
}

// Save saves the Reply to the database.
func (r *Reply) Save(db XODB) error {
	if r.Exists() {
		return r.Update(db)
	}

	return r.Insert(db)
}

// Upsert performs an upsert for Reply.
//
// NOTE: PostgreSQL 9.5+ only
func (r *Reply) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.replies (` +
		`snowflake, created_at, deleted_at, author_id, body, parent_id, topic_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (snowflake) DO UPDATE SET (` +
		`snowflake, created_at, deleted_at, author_id, body, parent_id, topic_id` +
		`) = (` +
		`EXCLUDED.snowflake, EXCLUDED.created_at, EXCLUDED.deleted_at, EXCLUDED.author_id, EXCLUDED.body, EXCLUDED.parent_id, EXCLUDED.topic_id` +
		`)`

	// run query
	XOLog(sqlstr, r.Snowflake, r.CreatedAt, r.DeletedAt, r.AuthorID, r.Body, r.ParentID, r.TopicID)
	_, err = db.Exec(sqlstr, r.Snowflake, r.CreatedAt, r.DeletedAt, r.AuthorID, r.Body, r.ParentID, r.TopicID)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Delete deletes the Reply from the database.
func (r *Reply) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return nil
	}

	// if deleted, bail
	if r._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.replies WHERE snowflake = $1`

	// run query
	XOLog(sqlstr, r.Snowflake)
	_, err = db.Exec(sqlstr, r.Snowflake)
	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

	return nil
}

// User returns the User associated with the Reply's AuthorID (author_id).
//
// Generated from foreign key 'replies_author_id_fkey'.
func (r *Reply) User(db XODB) (*User, error) {
	return UserBySnowflake(db, r.AuthorID.Int64)
}

// Reply returns the Reply associated with the Reply's ParentID (parent_id).
//
// Generated from foreign key 'replies_parent_id_fkey'.
func (r *Reply) Reply(db XODB) (*Reply, error) {
	return ReplyBySnowflake(db, r.ParentID.Int64)
}

// Topic returns the Topic associated with the Reply's TopicID (topic_id).
//
// Generated from foreign key 'replies_topic_id_fkey'.
func (r *Reply) Topic(db XODB) (*Topic, error) {
	return TopicBySnowflake(db, r.TopicID)
}

// RepliesByAuthorID retrieves a row from 'public.replies' as a Reply.
//
// Generated from index 'replies_author_index'.
func RepliesByAuthorID(db XODB, authorID sql.NullInt64) ([]*Reply, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`snowflake, created_at, deleted_at, author_id, body, parent_id, topic_id ` +
		`FROM public.replies ` +
		`WHERE author_id = $1`

	// run query
	XOLog(sqlstr, authorID)
	q, err := db.Query(sqlstr, authorID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Reply{}
	for q.Next() {
		r := Reply{
			_exists: true,
		}

		// scan
		err = q.Scan(&r.Snowflake, &r.CreatedAt, &r.DeletedAt, &r.AuthorID, &r.Body, &r.ParentID, &r.TopicID)
		if err != nil {
			return nil, err
		}

		res = append(res, &r)
	}

	return res, nil
}

// RepliesByParentID retrieves a row from 'public.replies' as a Reply.
//
// Generated from index 'replies_parent_index'.
func RepliesByParentID(db XODB, parentID sql.NullInt64) ([]*Reply, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`snowflake, created_at, deleted_at, author_id, body, parent_id, topic_id ` +
		`FROM public.replies ` +
		`WHERE parent_id = $1`

	// run query
	XOLog(sqlstr, parentID)
	q, err := db.Query(sqlstr, parentID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Reply{}
	for q.Next() {
		r := Reply{
			_exists: true,
		}

		// scan
		err = q.Scan(&r.Snowflake, &r.CreatedAt, &r.DeletedAt, &r.AuthorID, &r.Body, &r.ParentID, &r.TopicID)
		if err != nil {
			return nil, err
		}

		res = append(res, &r)
	}

	return res, nil
}

// ReplyBySnowflake retrieves a row from 'public.replies' as a Reply.
//
// Generated from index 'replies_pkey'.
func ReplyBySnowflake(db XODB, snowflake int64) (*Reply, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`snowflake, created_at, deleted_at, author_id, body, parent_id, topic_id ` +
		`FROM public.replies ` +
		`WHERE snowflake = $1`

	// run query
	XOLog(sqlstr, snowflake)
	r := Reply{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, snowflake).Scan(&r.Snowflake, &r.CreatedAt, &r.DeletedAt, &r.AuthorID, &r.Body, &r.ParentID, &r.TopicID)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// RepliesByTopicID retrieves a row from 'public.replies' as a Reply.
//
// Generated from index 'replies_topic_index'.
func RepliesByTopicID(db XODB, topicID int64) ([]*Reply, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`snowflake, created_at, deleted_at, author_id, body, parent_id, topic_id ` +
		`FROM public.replies ` +
		`WHERE topic_id = $1`

	// run query
	XOLog(sqlstr, topicID)
	q, err := db.Query(sqlstr, topicID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Reply{}
	for q.Next() {
		r := Reply{
			_exists: true,
		}

		// scan
		err = q.Scan(&r.Snowflake, &r.CreatedAt, &r.DeletedAt, &r.AuthorID, &r.Body, &r.ParentID, &r.TopicID)
		if err != nil {
			return nil, err
		}

		res = append(res, &r)
	}

	return res, nil
}

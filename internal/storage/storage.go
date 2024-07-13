package storage

import "errors"

// TODO rewrite on Psql

var (
	ErrURLExists   = errors.New("url/alias already exists in db")
	ErrURLNotFound = errors.New("url not found")
)

// ErrURLExists General db error processing function
//func ErrURLExists() error {
//	return fmt.Errorf("error from sqlite3: alias already exists in db")
//}

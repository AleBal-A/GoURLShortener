package storage

import "errors"

// TODO rewrite on Psql

var (
	ErrURLExists  = "url/alias already exists in db"
	ErrURLExists2 = errors.New("url/alias already exists in db")
)

// ErrURLExists General db error processing function
//func ErrURLExists() error {
//	return fmt.Errorf("error from sqlite3: alias already exists in db")
//}

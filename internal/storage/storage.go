package storage

import "errors"

// TODO rewrite on Psql

var (
	ErrURLExists   = errors.New("url/alias already exists in db")
	ErrURLNotFound = errors.New("url not found")
)

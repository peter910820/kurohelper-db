package kurohelperdb

import "errors"

var (
	ErrParameterNotFound = errors.New("kurohelperdb: parameter not found")
	ErrNoRowsAffected    = errors.New("kurohelperdb: No rows affected")
)

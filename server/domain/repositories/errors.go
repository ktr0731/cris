package repositories

import "errors"

var (
	ErrNotFound = errors.New("specified entity not found")
)

package storage

import (
	"errors"
)

var (
	ErrBrandExists   = errors.New("URL exists")
	ErrBrandNotFound = errors.New("brand not found")
)

type Storage interface {
	CloseDB() error
	GetName() string
}

package storage

import (
	"errors"
)

var (
	ErrBrandExists       = errors.New("brand exists")
	ErrBrandNotFound     = errors.New("brand not found")
	ErrVehicleTypeExists = errors.New("vehicle type exists")
	ErrBodyStyleExists   = errors.New("body style exists")
	ErrBodyStyleNotFound = errors.New("body style not found")
)

type Storage interface {
	CloseDB() error
	GetName() string
}

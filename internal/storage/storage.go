package storage

import (
	"errors"
)

var (
	ErrBrandExists         = errors.New("brand exists")
	ErrBrandNotFound       = errors.New("brand not found")
	ErrVehicleTypeExists   = errors.New("vehicle type exists")
	ErrVehicleTypeNotFound = errors.New("vehicle type not found")
)

type Storage interface {
	CloseDB() error
	GetName() string
}

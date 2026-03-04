package storage

import (
	"errors"
)

var (
	ErrBrandExists        = errors.New("brand exists")
	ErrVehicleExists      = errors.New("vehicle exists")
	ErrBrandNotFound      = errors.New("brand not found")
	ErrVehicleNotFound    = errors.New("vehicle not found")
	ErrBodyStyleExists    = errors.New("body style exists")
	ErrBodyStyleNotFound  = errors.New("body style not found")
	ErrCategoryNotFound   = errors.New("category not found")
	ErrDriverTypeNotFound = errors.New("driver type not found")
	ErrGearboxNotFound    = errors.New("gearbox not found")
)

type Storage interface {
	CloseDB() error
	GetName() string
}

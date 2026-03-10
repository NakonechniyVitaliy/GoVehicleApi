package _errors

import "errors"

var (
	ErrVehicleExists   = errors.New("vehicle exists")
	ErrVehicleNotFound = errors.New("vehicle not found")
)

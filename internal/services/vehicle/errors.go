package vehicle

import "errors"

var (
	ErrVehicleNotFound = errors.New("vehicle not found")
	ErrUpdateVehicle   = errors.New("failed to update vehicle")
	ErrVehicleExists   = errors.New("vehicle already exists")
	ErrSaveVehicle     = errors.New("failed to save vehicle")
	ErrGetVehicle      = errors.New("failed to get vehicle")
	ErrGetVehicles     = errors.New("failed to get vehicles")
)

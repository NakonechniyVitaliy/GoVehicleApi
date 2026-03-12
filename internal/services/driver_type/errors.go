package driver_type

import "errors"

var (
	ErrGetDriverTypes     = errors.New("failed to get driver types")
	ErrDecodeDriverTypes  = errors.New("failed to decode driver types")
	ErrDriverTypesFetch   = errors.New("failed to fetch driver types")
	ErrRefreshDriverTypes = errors.New("failed to refresh driver types")
)

package gearbox

import "errors"

var (
	ErrGetGearboxes     = errors.New("failed to get gearboxes")
	ErrDecodeGearboxes  = errors.New("failed to decode gearboxes")
	ErrGearboxesFetch   = errors.New("failed to fetch gearboxes")
	ErrRefreshGearboxes = errors.New("failed to refresh gearboxes")
)

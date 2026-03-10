package gearboxes

import "errors"

var (
	ErrGearboxesFetch  = errors.New("autoRia gearboxes fetch error")
	ErrDecodeGearboxes = errors.New("failed to decode gearboxes from autoRia")
)

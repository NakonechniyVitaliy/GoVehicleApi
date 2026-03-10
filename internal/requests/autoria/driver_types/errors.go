package autoria

import "errors"

var (
	ErrDriverTypesFetch  = errors.New("driver types fetch error")
	ErrDecodeDriverTypes = errors.New("autoRia failed to decode driver types from autoRia")
)

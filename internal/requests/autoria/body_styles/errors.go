package body_styles

import "errors"

var (
	ErrBodyStylesFetch  = errors.New("autoRia body styles fetch error")
	ErrDecodeBodyStyles = errors.New("failed to decode body styles from autoRia")
)

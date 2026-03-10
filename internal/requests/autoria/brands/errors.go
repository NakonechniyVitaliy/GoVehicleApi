package brands

import "errors"

var (
	ErrBrandsFetch  = errors.New("autoRia brands fetch error")
	ErrDecodeBrands = errors.New("autoRia failed to decode brands from autoRia")
)

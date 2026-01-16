package storage

import "errors"

var (
	ErrBrandExists = errors.New("URL exists")
)

type Storage interface {
	NewBrand() string
	RefreshBrands() string
}

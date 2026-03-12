package brand

import "errors"

var (
	ErrBrandNotFound = errors.New("brand not found")
	ErrUpdateBrand   = errors.New("failed to update brand")
	ErrBrandExists   = errors.New("brand already exists")
	ErrSaveBrand     = errors.New("failed to save brand")
	ErrGetBrand      = errors.New("failed to get brand")
	ErrGetBrands     = errors.New("failed to get brands")
	ErrDecodeBrands  = errors.New("failed to decode brands")
	ErrBrandsFetch   = errors.New("failed to fetch brands")
	ErrRefreshBrands = errors.New("failed to refresh brands")
)

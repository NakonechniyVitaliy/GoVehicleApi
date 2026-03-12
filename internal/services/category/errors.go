package category

import "errors"

var (
	ErrGetCategories     = errors.New("failed to get brands")
	ErrDecodeCategories  = errors.New("failed to decode brands")
	ErrCategoriesFetch   = errors.New("failed to fetch brands")
	ErrRefreshCategories = errors.New("failed to refresh brands")
)

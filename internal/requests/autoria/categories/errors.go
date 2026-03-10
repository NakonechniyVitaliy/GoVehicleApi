package categories

import "errors"

var (
	ErrCategoriesFetch  = errors.New("categories fetch error")
	ErrDecodeCategories = errors.New("autoRia failed to decode categories from autoRia")
)

package body_style

import "errors"

var (
	ErrBodyStyleNotFound = errors.New("body style not found")
	ErrUpdateBodyStyle   = errors.New("failed to update body style")
	ErrBodyStyleExists   = errors.New("body style already exists")
	ErrSaveBodyStyle     = errors.New("failed to save body style")
	ErrGetBodyStyle      = errors.New("failed to get body style")
	ErrGetBodyStyles     = errors.New("failed to get body styles")
	ErrDecodeBodyStyles  = errors.New("failed to decode body styles")
	ErrBodyStylesFetch   = errors.New("failed to fetch body styles")
	ErrRefreshBodyStyles = errors.New("failed to refresh body styles")
)

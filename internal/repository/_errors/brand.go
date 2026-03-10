package _errors

import "errors"

var (
	ErrBrandNotFound = errors.New("brand not found")
	ErrBrandExists   = errors.New("brand exists")
)

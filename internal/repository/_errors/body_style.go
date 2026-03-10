package _errors

import "errors"

var (
	ErrBodyStyleExists   = errors.New("body style exists")
	ErrBodyStyleNotFound = errors.New("body style not found")
)

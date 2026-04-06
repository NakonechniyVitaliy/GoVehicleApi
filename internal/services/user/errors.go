package user

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrIncorrectLoginOrPass = errors.New("incorrect login or password")
	ErrUserExists           = errors.New("user already exists")
	ErrSaveUser             = errors.New("failed to save user")
	ErrGetUser              = errors.New("failed to get user")
	ErrSignIn               = errors.New("failed to sign in")
	ErrComparePass          = errors.New("failed to compare passwords")
)

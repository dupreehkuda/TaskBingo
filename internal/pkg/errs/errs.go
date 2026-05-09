package errs

import "errors"

var (
	ErrEmptyRequest     = errors.New("empty request")
	ErrCredentialsInUse = errors.New("credentials already in use")
	ErrWrongCredentials = errors.New("wrong credentials")
	ErrNotFound         = errors.New("not found")
)

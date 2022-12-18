package customErrors

import "errors"

var (
	ErrCredentialsInUse = errors.New("username already in use")
	ErrWrongCredentials = errors.New("no such user or wrong password")
	ErrNoSuchPack       = errors.New("there is no such task pack")
)

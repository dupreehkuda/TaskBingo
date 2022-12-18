package customErrors

import "errors"

var (
	ErrNoSuchPack = errors.New("there is no such task pack")
)

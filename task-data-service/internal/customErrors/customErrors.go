package customErrors

import "errors"

var (
	ErrNoSuchPack        = errors.New("there is no such task pack")
	ErrPackAlreadyExists = errors.New("the pack with this name already exists")
)

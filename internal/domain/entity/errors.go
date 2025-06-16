package entity

import "errors"

var (
	ErrNotEnoughReserved = errors.New("not enough reserved stock")
	ErrNotEnoughStock    = errors.New("not enough available stock")
	ErrItemNotFound      = errors.New("item not found")
)

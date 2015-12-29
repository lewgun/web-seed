package errutil

import (
	"errors"
)

var (
	ErrIllegalParam  = errors.New("illegal parameter(s)")
	ErrMissingParams = errors.New("missing parameter(s)")
	ErrNotFound           = errors.New("not found")

)

package schedule

import (
	"errors"
)

var (
	ErrInvalidLocation   = errors.New("invalid location")
	ErrZeroNumberOfTimes = errors.New("number of times must be greater than 0")
)

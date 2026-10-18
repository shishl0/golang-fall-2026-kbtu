package task

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrNotFound          = errors.New("task not found")
	ErrInvalidTransition = errors.New("invalid state transition")
)

var ErrEmptyPassword = fmt.Errorf("%w: password is empty", ErrInvalidInput)

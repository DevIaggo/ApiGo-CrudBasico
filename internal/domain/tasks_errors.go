package domain

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrTaskNotFound    = errors.New("task not found")
	ErrTaskAlreadyDone = errors.New("task already done")
)

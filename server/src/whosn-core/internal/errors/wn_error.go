package errors

import (
	"errors"
	"fmt"
)

type WnError struct {
	StatusCode int

	Err error
}

func (e WnError) Error() string {
	return fmt.Sprintf("status %d: err %v", e.StatusCode, e.Err)
}

func NewError(status int, msg string) error {
	return &WnError{
		StatusCode: status,
		Err:        errors.New(msg),
	}
}

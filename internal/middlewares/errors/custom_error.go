package errors

import (
	"fmt"
)

type AppError struct {
	Cause string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("cause: %s",
		e.Cause)
}

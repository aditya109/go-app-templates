package errors

import (
	"fmt"
	"time"
)

type AppError struct {
	When  time.Time
	Cause string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.Cause)
}

/**
&AppError{
	time.Now(),
	"it didn't work",
}
*/
package common

import "fmt"

type ValidationError struct {
	Missing string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("missing fields: %s", e.Missing)
}

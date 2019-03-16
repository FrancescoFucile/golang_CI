package ci_test_1

import "fmt"

// MathError: mathematical error.
type MathError struct {
	errType string
}

func (e *MathError) Error() string {
	return fmt.Sprintf("mathematical error: %s", e.errType)
}

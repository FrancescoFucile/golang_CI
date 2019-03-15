package ci_test_1

import "fmt"

func Div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, mathError(fmt.Errorf("divide by zero"))
	}
	return a / b, nil
}

type mathError error

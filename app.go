package ci_test_1

import "fmt"

func Div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by 0")
	}
	return a / b, nil
}

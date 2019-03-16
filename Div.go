package ci_test_1

// API

// Div divides.
func Div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, &MathError{"divide by zero"}
	}
	return a / b, nil
}

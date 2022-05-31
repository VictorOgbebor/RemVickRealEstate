package bicmath

import "errors"

func Add(x, y int) int {
	return x + y
}

func Div(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannotDivide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

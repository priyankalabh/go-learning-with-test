package calc

import "errors"

func Division(dividend, divisor float32) (float32, error) {
	if divisor == 0 {
		return 0.0, errors.New("Divide by 0")
	}
	result := dividend / divisor
	return result, nil
}

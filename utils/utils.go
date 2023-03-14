package utils

import "errors"

func Min(number1, number2 int) (int, error) {

	if number1 > number2 {
		return number2, nil
	}
	if number1 == number2 {
		return 0, errors.New("numbers are Equal")
	}
	return number1, nil
}

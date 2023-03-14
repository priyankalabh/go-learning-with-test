package mystring

import "errors"

func Greetmsg(msg, name string) (string, error) {
	if len(msg) == 0 {
		return "", errors.New("msg required")
	}
	if len(name) == 0 {
		return "", errors.New("name required")
	}
	greetings := "Good" + " " + msg + ", " + name + "!"
	return greetings, nil
}

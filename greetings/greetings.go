package greetings

import (
	"errors"
	"fmt"
)

// Hello returns greetings for the named person
func Hello(name string) ( string, error ) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %s. Welcome!!!", name)
	return message, nil
}

package utils

import (
	"fmt"
)

func ReadInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}
	return input, nil
}
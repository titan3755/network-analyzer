package utils

import (
	"os"
	"errors"
)

func FileEmptyCheck(fileName string) (bool, error) {
	// check if file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false, errors.New("file does not exist")
	}
	// check if file is empty
	dat, err := os.ReadFile(fileName)
	if err != nil {
		return false, err
	}
	if string(dat) == "" {
		return true, nil
	}
	return false, nil
}

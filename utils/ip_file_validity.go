package utils

import (
	"fmt"
	"os"
	"strings"
)

func CheckIPFileFormatValidity(fileLocation string) error {
	// check if file exists
	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	// check if file is in correct format
	dat, err := os.ReadFile(fileLocation)
	if err != nil {
		return err
	}
	if string(dat) == "" {
		return fmt.Errorf("file is empty")
	}
	splitted := strings.Split(string(dat), ";")
	for _, line := range splitted {
		if string(line) != "" && !CheckIfValidIPv4(line) {
			return fmt.Errorf("invalid ip address")
		}
	}
	return nil
}
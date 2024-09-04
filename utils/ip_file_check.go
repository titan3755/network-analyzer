package utils

import (
	"os"
	"strings"
)

func CheckIfIPAlreadyInFile(ip string, fileLocation string) (bool, error) {
	// open and read file at location
	dat, err := os.ReadFile(fileLocation)
	if err != nil {
		return false, err
	}
	// check if ip is in file
	for _, line := range string(dat) {
		if strings.Contains(string(line), ip) {
			return true, nil
		}
	}
	return false, nil
}
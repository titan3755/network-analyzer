package utils

import (
	"errors"
	"os"
	"strings"
)

// IpFileReader is a function that reads a file containing a list of IP addresses and returns a list of IP addresses

func IpFileReader(fileName string) ([]string, error) {
	// check if file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil, errors.New("file does not exist")
	}
	// check if file is empty
	dat, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if string(dat) == "" {
		return nil, errors.New("file is empty")
	}
	// check file format validity
	errn := CheckIPFileFormatValidity(fileName)
	if errn != nil {
		return nil, errn
	}
	// remove duplicate ip from the file
	errev := RemoveDuplicateIPFromFile(fileName)
	if errev != nil {
		return nil, errev
	}
	// read file
	data, errfr := os.ReadFile(fileName)
	if errfr != nil {
		return nil, errfr
	}
	var finalList []string
	for _, line := range strings.Split(string(data), ";") {
		if line != "" {
			finalList = append(finalList, strings.TrimSpace(line))
		}
	}
	return finalList, nil
}

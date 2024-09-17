package utils

import (
	"os"
	"errors"
	"strings"
)

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
	var final_list []string
	for _, line := range strings.Split(string(data), ";") {
		if line != "" {
			final_list = append(final_list, strings.TrimSpace(line))
		}
	}
	return final_list, nil
}
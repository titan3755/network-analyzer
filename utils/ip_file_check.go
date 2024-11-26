package utils

import (
	"fmt"
	"os"
	"strings"
)

// this function checks if the ip is already in the file

func CheckIfIPAlreadyInFile(ip string, fileLocation string) (bool, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	var ips []string
	var temp string
	for {
		_, err := fmt.Fscan(file, &temp)
		if err != nil {
			break
		}
		ips = append(ips, temp)
	}
	for _, i := range ips {
		if strings.Contains(i, ip) {
			return true, nil
		}
	}
	return false, nil
}

// this function removes duplicate ip addresses from the file

func RemoveDuplicateIPFromFile(filePath string) error {
	// check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	// check if file is in correct format
	errn := CheckIPFileFormatValidity(filePath)
	if errn != nil {
		return errn
	}
	// open and read file at location
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	// detect duplicate ip and make list out of them
	var newDat string
	for _, line := range strings.Split(string(dat), ";") {
		if strings.Contains(newDat, line) {
			continue
		}
		newDat += line + ";"
	}
	// write new data to file
	file, er := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if er != nil {
		return er
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	_, err = file.WriteString(newDat)
	if err != nil {
		return err
	}
	return nil
}

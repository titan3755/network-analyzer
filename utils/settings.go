package utils

import (
	"os"
	"fmt"
	"strings"
)

// SetSettings is a function that sets a setting in the settings.prp file

func SetSettings(setting string, property string) error {
	// open and read file at location
	file, err := os.OpenFile("settings.prp", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(setting + "=" + property + "\n")
	if err != nil {
		return err
	}
	return nil
}

// GetSettings is a function that gets a setting from the settings.prp file

func GetSettings(setting string) (string, error) {
	// open and read file at location
	file, err := os.Open("settings.prp")
	if err != nil {
		return "", err
	}
	defer file.Close()
	var temp string
	for {
		_, err := fmt.Fscan(file, &temp)
		if err != nil {
			break
		}
		if strings.Contains(temp, setting) {
			return strings.Split(temp, "=")[1], nil
		}
	}
	return "", fmt.Errorf("setting not found")
}

// ReadSettings is a function that reads the settings from the settings.prp file

func ReadSettings(fileName string) map[string]string {
	// open and read file at location
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()
	var temp string
	settings := make(map[string]string)
	for {
		_, err := fmt.Fscan(file, &temp)
		if err != nil {
			break
		}
		settings[strings.Split(temp, "=")[0]] = strings.Split(temp, "=")[1]
	}
	return settings
}
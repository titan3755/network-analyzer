package utils

import (
	"errors"
	"netzer/data"
	"os"
)

// this function checks if the settings file exists in the same directory and if it does not exist, it creates it
// if it does exist, it does nothing

func SettingsFile() error {
	if _, err := os.Stat(data.SettingsFileName); errors.Is(err, os.ErrNotExist) {
		_, errr := os.Create(data.SettingsFileName)
		if errr != nil {
			return errr
		}
	} else if err != nil {
		return err
	}
	return nil
}

package utils

import (
	"os"
	"errors"
)

// this function checks if the settings file exists in the same directory and if it does not exist, it creates it
// if it does exist, it does nothing

func SettingsFile() error {
	if _, err := os.Stat("settings.prp"); errors.Is(err, os.ErrNotExist) {
		_, errr := os.Create("settings.prp")
		if errr != nil {
			return errr
		}
	} else if err != nil {
		return err
	}
	return nil
}
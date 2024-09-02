package utils

import (
	"os"
	"errors"
)

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
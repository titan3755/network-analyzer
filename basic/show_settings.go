package basic

import (
	"github.com/urfave/cli/v2"
	"os"
	"github.com/pterm/pterm"
	"netzer/utils"
	"fmt"
	"errors"
)

// this function reads the settings.prp file which exists in the same directory and then it
// prints the settings to the console or terminal

func ShowSettingsMain(c *cli.Context) error {
	utils.BasicIntro()
	// check if settings file exists
	if _, err := os.Stat("settings.prp"); errors.Is(err, os.ErrNotExist) {
		pterm.Error.Println("settings file does not exist")
		return fmt.Errorf("settings file does not exist")
	} else if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	// check if settings file is empty
	empF, ern := utils.FileEmptyCheck("settings.prp")
	if ern != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", ern))
		return ern
	}
	if empF {
		pterm.Info.Println("settings file is empty")
		return nil
	}
	// read settings file
	sett := utils.ReadSettings("settings.prp")
	if sett == nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", errors.New("could not read settings file")))
		return fmt.Errorf("could not read settings file")
	}
	// print settings
	pterm.Success.Println("Successfully read settings file; output -->")
	for key, value := range sett {
		pterm.Info.Println(fmt.Sprintf("%v: %v", key, value))
	}
	return nil
}
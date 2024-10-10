package basic

import (
	"fmt"
	"netzer/data"
	"netzer/utils"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func SetSettingsMain(c *cli.Context) error {
	utils.BasicIntro()
	settingFile := "settings.prp"
	fmt.Print("\n")
	file := utils.ReadSettings(settingFile)
	if file == nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", fmt.Errorf("something wrong with settings file or file empty")))
		return fmt.Errorf("could not read settings file")
	}
	pterm.Info.Println("Settings file read successfully. Output -->")
	for key, value := range file {
		pterm.Info.Println(fmt.Sprintf("%v: %v", key, value))
	}
	fmt.Print("\n")
	possibleSettings := data.AvailableSettings
	var optionsSettings []string
	optionsSettings = append(optionsSettings, possibleSettings...)
	pterm.Info.Println("Available settings:")
	fmt.Print("\n")
	selectedOptionSetting, _ := pterm.DefaultInteractiveSelect.WithOptions(optionsSettings).Show()
	pterm.Info.Println(fmt.Sprintf("Selected setting: %v", selectedOptionSetting))
	fmt.Print("\n")
	pterm.Info.Println("Enter the property for the setting:")
	fmt.Print("\n")
	property, erio := utils.ReadInput()
	if erio != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", erio))
		return erio
	}
	err := utils.SetSettings(selectedOptionSetting, property)
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("Successfully set setting")
	return nil
}
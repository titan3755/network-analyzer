package basic

import (
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
)

func WipeSettingsMain(_ *cli.Context) error {
	utils.BasicIntro()
	pterm.Info.Println("Wiping settings...")
	err := utils.WipeSettings()
	if err != nil {
		pterm.Error.Println("Error: ", err)
		return err
	}
	pterm.Success.Println("Successfully wiped settings")
	return nil
}

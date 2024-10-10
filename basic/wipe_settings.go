package basic

import (
	"github.com/urfave/cli/v2"
	"netzer/utils"
	pterm "github.com/pterm/pterm"
)

func WipeSettingsMain(c *cli.Context) error {
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
package basic

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/data"
	"netzer/utils"
)

// this function shows the version of the application

func ShowVersion(_ *cli.Context) error {
	utils.BasicIntro()
	pterm.Info.Println("Version: ", data.CurrentAppVersion)
	fmt.Print("\n")
	style := pterm.NewStyle(pterm.BgBlack, pterm.FgLightGreen, pterm.Bold)
	style.Println("Previous versions -->")
	fmt.Print("\n")
	for _, version := range data.PreviousAppVersionsInclLatestVersion {
		pterm.Info.Println(version)
	}
	fmt.Print("\n")
	return nil
}

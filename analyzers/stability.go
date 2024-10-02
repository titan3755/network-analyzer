package analyzers

import (
	"netzer/utils"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func StabilityAnalyzerMain(c *cli.Context) error {
	utils.AnalyzerIntro()
	var analyzingTime int = 20
	pterm.Info.Println("This analyzer will create a stability report for the network by pinging different servers and checking the packet losses.")
	pterm.Info.Println("The analyzer will run for", analyzingTime, "seconds.")
	return nil
}	
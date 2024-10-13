package analyzers

import (
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"strconv"
	"github.com/pterm/pterm"
)

func StabilityAnalyzerFullMain(c *cli.Context) error {
	utils.AnalyzerIntro()
	var analyzingTime int = 30
	if c.Args().First() != "" {
		timeToInt, err := strconv.Atoi(c.Args().First())
		if err != nil {
			pterm.Warning.Println("An error occurred while converting the time to integer. Using default time of 20 seconds ...")
		} else if timeToInt < 30 {
			pterm.Warning.Println("The time entered is less than 30 seconds. Using default time of 30 seconds ...")
		} else {
			analyzingTime = timeToInt
		}
	}
	pterm.Info.Println("This analyzer will create a complete network stability report by pinging different servers, checking the packet losses and simultaneously speed testing the network.")
	pterm.Info.Println("The speed testing is done to ensure that the network can hold up under load.")
	pterm.Info.Println("The analyzer will run for", analyzingTime, "seconds.")
	// to complete
	return nil
}		
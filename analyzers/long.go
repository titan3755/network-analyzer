package analyzers

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"strconv"
	"sync"
)

func StabilityAnalyzerLongMain(c *cli.Context) error {
	utils.AnalyzerIntro()
	var analyzingTime = 5 // the time is in minutes here
	if c.Args().First() != "" {
		timeToInt, err := strconv.Atoi(c.Args().First())
		if err != nil {
			pterm.Warning.Println(fmt.Sprintf("An error occurred while converting the time to integer. Using default time of %d minutes ...", analyzingTime))
		} else if timeToInt < analyzingTime {
			pterm.Warning.Println(fmt.Sprintf("The time entered is less than %d minutes. Using default time of %d minutes ...", analyzingTime, analyzingTime))
		} else {
			analyzingTime = timeToInt
		}
	}
	var wg sync.WaitGroup
	pterm.Info.Println("This analyzer will create a complete network stability report by pinging different servers, checking the packet losses and simultaneously speed testing the network")
	pterm.Info.Println("The speed testing is done to ensure that the network can hold up under load")
	pterm.Info.Println("This specific analyzer will run for a longer time in the background to ensure proper network stability testing")
	pterm.Info.Println("The data will be written to a file or the existing data in the file will be updated every 5 minutes")
	pterm.Info.Println("The analyzer will run for", analyzingTime, "minutes.")
	pterm.Info.Println("Starting the long stability test [background process] ...")
	pterm.Info.Println("Note that the total required time may be greater than the analyzing time due to the speed test")
	pterm.Info.Println("DO NOT CLOSE OR TERMINATE THIS PROCESS/WINDOW OR THE ANALYZER WILL STOP RUNNING")
	wg.Add(1)
	go utils.AnalyzerBackgroundProcess(analyzingTime, &wg)
	wg.Wait()
	return nil
}

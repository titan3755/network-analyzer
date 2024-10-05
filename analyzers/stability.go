package analyzers

import (
	"netzer/utils"
	"strconv"
	"netzer/data"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func StabilityAnalyzerMain(c *cli.Context) error {
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
	pterm.Info.Println("This analyzer will create a stability report for the network by pinging different servers and checking the packet losses.")
	pterm.Info.Println("The analyzer will run for", analyzingTime, "seconds.")
	var longIPList map[string][]string = make(map[string][]string)
	var errsa []error
	var ipList []string
	longIPList, errsa = utils.ConvertListOfHostsToIPs(data.StabilityTestAddrList)
	if len(errsa) > 0 {
		pterm.Error.Println("An error occurred while converting the list of hosts to IP addresses.")
		for _, err := range errsa {
			pterm.Error.Println(err)
		}
	}
	ipList = data.StabilityTestIPList
	var mergedIPList []string
	var spinnerOn bool = true
	mergedIPList = append(mergedIPList, ipList...)
	for _, ipl := range longIPList {
		mergedIPList = append(mergedIPList, ipl...)
	}
	pterm.Info.Println("The following IP addresses will be tested:")
	for _, ip := range mergedIPList {
		pterm.Info.Println(ip)
	}
	pterm.Info.Println("Starting the stability test ...")
	go func() {
		spnrInfo, _ := pterm.DefaultSpinner.Start("Performing stability test ...")
		for {
			if !spinnerOn {
				spnrInfo.Success("Stability test completed!")
				break
			}
			continue
		}
	}()
	ipMap, errMap := utils.ICMP_Ping_Concurrent(mergedIPList, analyzingTime)
	spinnerOn = false
	pterm.Info.Println("Generating results ...")
	utils.StatisticsTableCreatorForStabilityAnalyzer(ipMap, errMap)
	return nil
}	
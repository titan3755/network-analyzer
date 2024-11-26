package analyzers

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/data"
	"netzer/utils"
	"strconv"
	"time"
)

func StabilityAnalyzerFullMain(c *cli.Context) error {
	utils.AnalyzerIntro()
	var analyzingTime = 60
	if c.Args().First() != "" {
		timeToInt, err := strconv.Atoi(c.Args().First())
		if err != nil {
			pterm.Warning.Println("An error occurred while converting the time to integer. Using default time of 60 seconds ...")
		} else if timeToInt < 60 {
			pterm.Warning.Println("The time entered is less than 60 seconds. Using default time of 60 seconds ...")
		} else {
			analyzingTime = timeToInt
		}
	}
	pterm.Info.Println("This analyzer will create a complete network stability report by pinging different servers, checking the packet losses and simultaneously speed testing the network.")
	pterm.Info.Println("The speed testing is done to ensure that the network can hold up under load.")
	pterm.Info.Println("The analyzer will run for", analyzingTime, "seconds.")
	var longIPList = make(map[string][]string)
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
	var spinnerOn = true
	mergedIPList = append(mergedIPList, ipList...)
	for _, ipl := range longIPList {
		mergedIPList = append(mergedIPList, ipl...)
	}
	pterm.Info.Println("The following hosts will be tested:")
	for _, host := range data.StabilityTestAddrList {
		fmt.Printf("[%s]", host)
	}
	fmt.Print("\n\n")
	pterm.Info.Println("The following IP addresses will be tested:")
	for _, ip := range data.StabilityTestIPList {
		fmt.Printf("[%s]", ip)
	}
	fmt.Print("\n\n")
	pterm.Info.Println("Starting the full stability test ...")
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
	// speed test
	var speedTestData = make(map[string][][]string)
	go func() {
		speedTestData = utils.SpeedTestAll(analyzingTime)
	}()
	// stability test
	var stabilityTestData = make(map[string][][]string)
	var errMap map[string][]error
	go func() {
		stabilityTestData, errMap = utils.IcmpPingConcurrent(mergedIPList, analyzingTime)
		if len(errMap) > 0 {
			pterm.Error.Println("An error occurred while performing the stability test.")
			for key, value := range errMap {
				pterm.Error.Println(key)
				pterm.Error.Println(value)
			}
		}
	}()
	var timeStamp = time.Now().Unix()
	var stopTime = timeStamp + int64(analyzingTime)
	var timeRunOut = false
	// check if time run out
	go func() {
		for {
			if time.Now().Unix() >= stopTime {
				timeRunOut = true
				break
			}
			continue
		}
	}()
	// wait for both tests to complete
	for {
		if (len(speedTestData) > 0 && len(stabilityTestData) > 0) || timeRunOut {
			break
		}
		continue
	}
	// print the results
	pterm.Info.Println("Generating results ...")
	// table creator location
	// for now, just print out the raw data for both tests (debugging purposes)
	pterm.Info.Println("Speed test data:")
	for key, value := range speedTestData {
		pterm.Info.Println(key)
		for _, val := range value {
			pterm.Info.Println(val)
		}
	}
	pterm.Info.Println("Stability test data:")
	for key, value := range stabilityTestData {
		pterm.Info.Println(key)
		for _, val := range value {
			pterm.Info.Println(val)
		}
	}
	return nil
}

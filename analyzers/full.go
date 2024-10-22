package analyzers

import (
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"netzer/data"
	"strconv"
	"github.com/pterm/pterm"
	"time"
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
	pterm.Info.Println("The following hosts will be tested:")
	for _, host := range data.StabilityTestAddrList {
		pterm.Info.Printf("[%s]", host)
	}
	pterm.Info.Println("\nThe following IP addresses will be tested:")
	for _, ip := range data.StabilityTestIPList {
		pterm.Info.Printf("[%s]", ip)
	}
	pterm.Info.Println("\nStarting the full stability test ...")
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
	var speed_test_data map[string][][]string = make(map[string][][]string)
	go func() {
		speed_test_data = utils.SpeedTestAll(analyzingTime)
	}()
	// stability test
	var stability_test_data map[string][][]string = make(map[string][][]string)
	var errMap map[string][]error
	go func() {
		stability_test_data, errMap = utils.ICMP_Ping_Concurrent(mergedIPList, analyzingTime)
		if len(errMap) > 0 {
			pterm.Error.Println("An error occurred while performing the stability test.")
			for key, value := range errMap {
				pterm.Error.Println(key)
				pterm.Error.Println(value)
			}
		}
	}()
	var time_stamp = time.Now().Unix()
	var stop_time = time_stamp + int64(analyzingTime)
	var time_run_out bool = false
	// check if time run out
	go func() {
		for {
			if time.Now().Unix() >= stop_time {
				time_run_out = true
				break
			}
			continue
		}
	}()
	// wait for both tests to complete
	for {
		if (len(speed_test_data) > 0 && len(stability_test_data) > 0) || time_run_out {
			break
		}
		continue
	}
	// print the results
	pterm.Info.Println("Generating results ...")
	// table creator location
	// for now, just print out the raw data for both tests (debugging purposes)
	pterm.Info.Println("Speed test data:")
	for key, value := range speed_test_data {
		pterm.Info.Println(key)
		for _, val := range value {
			pterm.Info.Println(val)
		}
	}
	pterm.Info.Println("Stability test data:")
	for key, value := range stability_test_data {
		pterm.Info.Println(key)
		for _, val := range value {
			pterm.Info.Println(val)
		}
	}
	return nil
}		
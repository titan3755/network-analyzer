package utils

import (
	"fmt"
	"github.com/pterm/pterm"
	"netzer/data"
	"sync"
	t "time"
)

func AnalyzerBackgroundProcess(time int, wgs *sync.WaitGroup) {
	defer wgs.Done()
	var wg sync.WaitGroup
	pterm.Info.Println("Starting background process for analyzing data ...")
	var longIPList = make(map[string][]string)
	var errsa []error
	var ipList []string
	longIPList, errsa = ConvertListOfHostsToIPs(data.StabilityTestAddrList)
	if len(errsa) > 0 {
		pterm.Error.Println("An error occurred while converting the list of hosts to IP addresses.")
		for _, err := range errsa {
			pterm.Error.Println(err)
		}
	}
	ipList = data.StabilityTestIPList
	var mergedIPList []string
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

	// Declare the channel
	speedTestChan := make(chan map[string][][]string)

	// speed test background process
	var speedTestData = make(map[string][][]string)
	wg.Add(2)
	go func() {
		pterm.Info.Println("Starting speed test background process ...")
		defer wg.Done()
		var wgd sync.WaitGroup
		for i := 0; i < time; i++ {
			wgd.Add(1)
			go func() {
				defer wgd.Done()
				speedTestData = SpeedTestLongAnalyzer(speedTestChan)
			}()
			wgd.Wait()
		}
	}()

	// stability test background process
	var stabilityTestData = make(map[string][][]string)
	var errMap map[string][]error
	go func() {
		pterm.Info.Println("Starting stability test background process ...")
		defer wg.Done()
		var wgd sync.WaitGroup
		for i := 0; i < time; i++ {
			wgd.Add(1)
			go func() {
				defer wgd.Done()
				stabilityTestData, errMap = IcmpPingConcurrent(mergedIPList, 60)
				if len(errMap) > 0 {
					pterm.Error.Println("An error occurred while performing the stability test.")
					for key, value := range errMap {
						pterm.Error.Println(key)
						pterm.Error.Println(value)
					}
				}
			}()
			wgd.Wait()
		}
	}()

	// Receive data from the speed test channel
	go func() {
		for d := range speedTestChan {
			for key, value := range d {
				speedTestData[key] = value
			}
		}
	}()

	go func() {
		for i := 0; i < time; i++ {
			pterm.Info.Println("Writing data to file ...")
			// write data to file
			var stbSuccess = OutputAnalyzerDataToFile(stabilityTestData, data.StabilityTestDataFileName)
			var spdSuccess = OutputAnalyzerDataToFile(speedTestData, data.SpeedTestDataFileName)
			if !stbSuccess {
				pterm.Error.Println("An error occurred while writing the stability test data to file.")
			}
			if !spdSuccess {
				pterm.Error.Println("An error occurred while writing the speed test data to file.")
			}
			t.Sleep(1 * t.Minute)
		}
	}()
	wg.Wait()
	pterm.Info.Println("Written obtained data to files [speed_test_data.data, stability_test_data.data] ...")
}

// incomplete

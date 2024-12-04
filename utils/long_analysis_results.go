package utils

import (
	"fmt"
	"github.com/pterm/pterm"
)

func ShowAnalyzerSpeedTestData(data map[string][][]string) {
	pterm.Println("The following data was obtained from the speed test:")
	for key, value := range data {
		fmt.Print("\n")
		pterm.Info.Println("Speed Test Data for " + key + ":")
		for i, v := range value[0] {
			switch i {
			case 0:
				pterm.Info.Println("Hostname: " + v)
			case 1:
				pterm.Info.Println("Host Addr: " + v)
			case 2:
				pterm.Info.Println("Ping: " + v)
			case 3:
				pterm.Info.Println("Download Speed: " + v)
			case 4:
				pterm.Info.Println("Upload Speed: " + v)
			}
		}
	}
}

func ShowAnalyzerStabilityTestData(data map[string][][]string) {
	pterm.Info.Println("The following data was obtained from the stability test:")
	for key, value := range data {
		fmt.Print("\n")
		pterm.Info.Println("Stability Test Data for " + key + ":")
		if len(value) > 0 && len(value[0]) > 0 {
			for i, v := range value[0] {
				switch i {
				case 0:
					pterm.Info.Println("Hostname: " + v)
				case 1:
					pterm.Info.Println("Host Addr: " + v)
				case 2:
					pterm.Info.Println("Ping: " + v)
				case 3:
					pterm.Info.Println("Packets Sent: " + v)
				case 4:
					pterm.Info.Println("Packets Received: " + v)
				case 5:
					pterm.Info.Println("Packet Loss: " + v)
				case 6:
					pterm.Info.Println("Min RTT: " + v)
				case 7:
					pterm.Info.Println("Max RTT: " + v)
				case 8:
					pterm.Info.Println("Avg RTT: " + v)
				}
			}
		} else {
			pterm.Warning.Println("No data available for " + key)
		}
	}
}

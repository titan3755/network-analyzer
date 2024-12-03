package utils

import (
	"github.com/pterm/pterm"
)

func ShowAnalyzerSpeedTestData(data map[string][][]string) {
	pterm.Println("The following data was obtained from the speed test:")
	pterm.Info.Println(data)
}

func ShowAnalyzerStabilityTestData(data map[string][][]string) {
	pterm.Info.Println("The following data was obtained from the stability test:")
	pterm.Info.Println(data)
}

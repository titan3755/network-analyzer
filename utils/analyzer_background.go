package utils

import (
	"github.com/pterm/pterm"
	"sync"
	t "time"
)

func AnalyzerBackgroundProcess(time int, wg *sync.WaitGroup) {
	defer wg.Done()
	pterm.Info.Println("Starting background process for analyzing data ...")
	// speed test background process
	go func() {
		pterm.Info.Println("Starting speed test background process ...")
		// to do
	}()
	// stability test background process
	go func() {
		pterm.Info.Println("Starting stability test background process ...")
		// to do
	}()
	for i := 0; i < time; i++ {
		pterm.Info.Println("Analyzing data ...")
		t.Sleep(1 * t.Minute)
	}
	pterm.Info.Println("Data analysis complete ...")
}

// incomplete

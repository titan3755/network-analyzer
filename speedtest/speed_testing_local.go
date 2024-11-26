package speedtest

import (
	"fmt"
	"github.com/pterm/pterm"
	st "github.com/showwin/speedtest-go/speedtest"
	"github.com/urfave/cli/v2"
	"netzer/utils"
)

// SpeedTestLocalMain function is the main function for the local speed test

func SpeedTestLocalMain(c *cli.Context) error {
	utils.SpeedTestIntro()
	var spinnerOn = true
	pterm.Info.Println("Starting quick speed test to domestic server ... [powered by speedtest.net]")
	pterm.Info.Println("[Go API by showwin (https://github.com/showwin/speedtest-go)]")
	fmt.Print("\n\n")
	go func() {
		spnrInfo, _ := pterm.DefaultSpinner.Start("Performing speed test ...")
		for {
			if !spinnerOn {
				spnrInfo.Success("Speed test completed!")
				break
			}
			continue
		}
	}()
	var speedTester = st.New()
	serverLst, _ := speedTester.FetchServers()
	trgts, _ := serverLst.FindServer([]int{})
	for _, srvr := range trgts {
		err := srvr.PingTest(nil)
		if err != nil {
			return err
		}
		err = srvr.DownloadTest()
		if err != nil {
			return err
		}
		err = srvr.UploadTest()
		if err != nil {
			return err
		}
		pterm.Info.Printf("Latency: %s, Download: %s, Upload: %s\n", srvr.Latency, srvr.DLSpeed, srvr.ULSpeed)
		spinnerOn = false
		srvr.Context.Reset()
	}
	return nil
}

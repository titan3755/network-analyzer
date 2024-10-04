package speedtest

import (
	"netzer/utils"
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	st "github.com/showwin/speedtest-go/speedtest"
)

func SpeedTestGlobalMain(c *cli.Context) error {
	utils.SpeedTestIntro()
	var spinnerOn bool = true
	var speedTestingURL string = c.Args().First()
	var speed_tester = st.New()
	var mode bool
	pterm.Info.Println("Starting quick speed test to global server ... [powered by speedtest.net]")
	pterm.Info.Println("[Go API by showwin (https://github.com/showwin/speedtest-go)]")
	fmt.Print("\n\n")
	if speedTestingURL == "" {
		pterm.Info.Println("No URL provided. Performing speed test with respect to closest server ...")
		mode = false
	} else {
		pterm.Info.Println(fmt.Sprintf("Performing speed test with respect to server: %s ...", speedTestingURL))
		mode = true
	}
	// spinner code here -->
	go func() {
		spnrInfo, _ := pterm.DefaultSpinner.Start("Performing speed test (this may take a while) ...")
		for {
			if !spinnerOn {
				spnrInfo.Success("Speed test completed!")
				break
			}
			continue
		}
	}()
	// spinner code here <--
	if mode {
		pterm.Info.Printf("mode check: %v\n", mode)
		serverList, _ := speed_tester.FetchServers()
		fmt.Printf("serverList: %v\n", serverList.String())
		spinnerOn = false
		// to do: complete the code here
	} else {
		serverList, _ := speed_tester.FetchServers()
		trgts, _ := serverList.FindServer([]int{})
		for _, srvr := range trgts {
			srvr.PingTest(nil)
			srvr.DownloadTest()
			srvr.UploadTest()
			pterm.Info.Printf("Latency: %s, Download: %s, Upload: %s\n", srvr.Latency, srvr.DLSpeed, srvr.ULSpeed)
			spinnerOn = false
			srvr.Context.Reset()
		}
	}
	return nil
}

// to complete
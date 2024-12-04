package speedtest

import (
	"fmt"
	"github.com/pterm/pterm"
	st "github.com/showwin/speedtest-go/speedtest"
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"slices"
	"strconv"
)

// SpeedTestGlobalMain function is the main function for the global speed test

func SpeedTestGlobalMain(_ *cli.Context) error {
	utils.SpeedTestIntro()
	var spinnerOn = true
	var speedTester = st.New()
	pterm.Info.Println("Starting quick speed test to global server ... [powered by speedtest.net]")
	pterm.Info.Println("[Go API by showwin (https://github.com/showwin/speedtest-go)]")
	fmt.Print("\n\n")
	pterm.Info.Println("Select a region to test against -->")
	var options []string
	for item := range st.Locations {
		options = append(options, item)
	}
	selectedOptionRegion, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOptionRegion))
	fmt.Print("\n")
	speedTester.NewUserConfig(&st.UserConfig{Location: st.Locations[selectedOptionRegion]})
	pterm.Info.Println("Select a server to test against -->")
	serverList, _ := speedTester.FetchServers()
	var serverOptions []string
	var srvrIDLst []string
	for _, server := range serverList {
		serverOptions = append(serverOptions, server.Name)
		srvrIDLst = append(srvrIDLst, server.ID)
	}
	selectedOptionServer, _ := pterm.DefaultInteractiveSelect.WithOptions(serverOptions).Show()
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOptionServer))
	fmt.Print("\n")
	selectedSrvrID := srvrIDLst[slices.Index(serverOptions, selectedOptionServer)]
	srvrIDToInt, _ := strconv.Atoi(selectedSrvrID)
	trgt, _ := serverList.FindServer([]int{srvrIDToInt})
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
	for _, srv := range trgt {
		err := srv.PingTest(nil)
		if err != nil {
			return err
		}
		err = srv.DownloadTest()
		if err != nil {
			return err
		}
		err = srv.UploadTest()
		if err != nil {
			return err
		}
		pterm.Info.Printf("Latency: %s, Download: %s, Upload: %s\n", srv.Latency, srv.DLSpeed, srv.ULSpeed)
		spinnerOn = false
		srv.Context.Reset()
	}
	// to do: complete the code here
	return nil
}

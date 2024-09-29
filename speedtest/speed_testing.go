package speedtest

import (
	"fmt"
	"netzer/utils"
	"github.com/pterm/pterm"
	st "github.com/showwin/speedtest-go/speedtest"
	"github.com/urfave/cli/v2"
)

func SpeedTestMain(c *cli.Context) error {
	utils.SpeedTestIntro()
	pterm.Info.Println("Starting speed test ... [powered by speedtest.net]")
	pterm.Info.Println("[Go API by showwin (https://github.com/showwin/speedtest-go)]")
	fmt.Print("\n\n")
	var speed_tester = st.New()
	serverLst, _ := speed_tester.FetchServers()
	trgts, _ := serverLst.FindServer([]int{})
	for _, srvr := range trgts {
		srvr.PingTest(nil)
		srvr.DownloadTest()
		srvr.UploadTest()
		pterm.Info.Printf("Latency: %s, Download: %s, Upload: %s\n", srvr.Latency, srvr.DLSpeed, srvr.ULSpeed)
		srvr.Context.Reset()
	}
	return nil
}
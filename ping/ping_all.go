package ping

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"strconv"
	"time"
)

// this function pings all the IP addresses in the IP file (main_cmd_function)

func PingAllMain(c *cli.Context) error {
	utils.PingIntro()
	pterm.Info.Println("Reading settings file for IP file location ...")
	ipFile, err := utils.GetSettings("ip_file")
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Info.Println(fmt.Sprintf("IP file location: %v", ipFile))
	pterm.Info.Println("Trying to ping IP addresses in IP file ...")
	pterm.Info.Println("Reading IP file ...")
	ipList, errs := utils.IpFileReader(ipFile)
	if errs != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", errs))
		return errs
	}
	pterm.Info.Println("Listing IPs found in file ...")
	for no, ip := range ipList {
		pterm.Info.Println(fmt.Sprintf("%v. %v \n", no+1, ip))
	}
	pterm.Info.Println("Pinging IPs ...")
	fmt.Print("\n\n")
	var pingTimeInSeconds = 20
	if c.Args().Get(0) != "" {
		pingTimeInSecondsInt, erhsh := strconv.Atoi(c.Args().Get(0))
		if erhsh != nil {
			pingTimeInSeconds = 20
			pterm.Error.Println("Error: Invalid time specified, using default time of 20 seconds")
		} else {
			if pingTimeInSecondsInt < 5 || pingTimeInSecondsInt > 120 {
				pingTimeInSeconds = 20
				pterm.Error.Println("Error: Time specified is out of range, using default time of 20 seconds")
			} else {
				pingTimeInSeconds = pingTimeInSecondsInt
				pterm.Info.Println(fmt.Sprintf("Pinging for %d seconds ...", pingTimeInSeconds))
			}
		}
	} else {
		pingTimeInSeconds = 20
		pterm.Info.Println("Pinging for 20 seconds ...")
	}
	// progress bar area ---
	successSpinner, _ := pterm.DefaultSpinner.Start("Pinging ...")
	// --- progress bar area
	ipMapMain, errors := utils.IcmpPingConcurrent(ipList, pingTimeInSeconds)
	for _, dt := range errors {
		if len(dt) > 0 {
			successSpinner.Warning("Some errors occurred - displaying valid data ...")
		} else {
			successSpinner.Success("Pinging complete ...")
		}
	}
	fmt.Print("\n")
	pterm.Success.Println("Pinging complete ...")
	fmt.Print("\n")
	statSpinner, _ := pterm.DefaultSpinner.Start("Creating statistics table ...")
	time.Sleep(2 * time.Second)
	statSpinner.Success("Statistics table created ...")
	fmt.Print("\n")
	utils.ResetTerminal()
	// --- statistics table area
	utils.StatisticsTableCreatorForPingAll(ipMapMain, errors)
	// statistics table area ---
	return nil
}

package ping

import (
	"fmt"
	"netzer/utils"
	"strconv"
	"time"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

// this function pings all the IP addresses in the IP file (main_cmd_function)

func PingAllMain(c *cli.Context) error {
	utils.PingIntro()
	pterm.Info.Println("Reading settings file for IP file location ...")
	ip_file, err := utils.GetSettings("ip_file")
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Info.Println(fmt.Sprintf("IP file location: %v", ip_file))
	pterm.Info.Println("Trying to ping IP addresses in IP file ...")
	pterm.Info.Println("Reading IP file ...")
	ip_list, errs := utils.IpFileReader(ip_file)
	if errs != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", errs))
		return errs
	}
	pterm.Info.Println("Listing IPs found in file ...")
	for no, ip := range ip_list {
		pterm.Info.Println(fmt.Sprintf("%v. %v \n", no+1, ip))
	}
	pterm.Info.Println("Pinging IPs ...")
	fmt.Print("\n\n")
	var ping_time_in_seconds int = 20
	if c.Args().Get(0) != "" {
		ping_time_in_seconds_int, erhsh := strconv.Atoi(c.Args().Get(0))
		if erhsh != nil {
			ping_time_in_seconds = 20
			pterm.Error.Println("Error: Invalid time specified, using default time of 20 seconds")
		} else {
			if ping_time_in_seconds_int < 5 || ping_time_in_seconds_int > 120 {
				ping_time_in_seconds = 20
				pterm.Error.Println("Error: Time specified is out of range, using default time of 20 seconds")
			} else {
				ping_time_in_seconds = ping_time_in_seconds_int
				pterm.Info.Println(fmt.Sprintf("Pinging for %d seconds ...", ping_time_in_seconds))
			}
		}
	} else {
		ping_time_in_seconds = 20
		pterm.Info.Println("Pinging for 20 seconds ...")
	}
	// progress bar area ---
	successSpinner, _ := pterm.DefaultSpinner.Start("Pinging ...")
	// --- progress bar area
	ip_map_main, errors := utils.ICMP_Ping_Concurrent(ip_list, ping_time_in_seconds)
	for _, dt := range errors {
		if len(dt) > 0{
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
	utils.StatisticsTableCreatorForPingAll(ip_map_main, errors)
	// statistics table area ---
	return nil
}
package ping

import (
	"fmt"
	"netzer/utils"
	"slices"
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
		return err
	}
	pterm.Info.Println(fmt.Sprintf("IP file location: %v", ip_file))
	pterm.Info.Println("Trying to ping IP addresses in IP file ...")
	pterm.Info.Println("Reading IP file ...")
	ip_list, errs := utils.IpFileReader(ip_file)
	if errs != nil {
		return errs
	}
	pterm.Info.Println("Listing IPs found in file ...")
	for no, ip := range ip_list {
		pterm.Info.Println(fmt.Sprintf("%v. %v \n", no+1, ip))
	}
	pterm.Info.Println("Pinging IPs ...")
	fmt.Print("\n\n")
	area, _ := pterm.DefaultArea.Start()
	comms := make(chan []string)
	defer area.Stop()
	for {
		// use icmp_ping_concurrent to ping all the IPs
		// use a channel to communicate with the function
		// print the results
		// stop the function when all IPs are pinged
		for _, ip := range ip_list {
			go utils.ICMP_Ping_Concurrent(ip, comms)
		}
		var areaUpdateStr string
		var ip_up_list []string
		msg := <- comms
		go func() {
			go func () {
				for {
					time.Sleep(time.Second * 2)
					if len(ip_up_list) == len(ip_list) {
						areaUpdateStr = ""
						ip_up_list = []string{}
					}
				}
			}()
			for {
				msg = <- comms
				if !slices.Contains(ip_up_list, msg[6]) && slices.Contains(ip_list, msg[6]) {
					// add to ip_up_list and update areaUpdateStr
					ip_up_list = append(ip_up_list, msg[6])
					highest, _ := strconv.Atoi(msg[2])
					lowest, _ := strconv.Atoi(msg[3])
					areaUpdateStr += pterm.Info.Sprintf("IP: %s\nLatency: %s\nAverage: %s\nHighest: %dms\nLowest: %dms\n\n", msg[6], msg[0], msg[1], highest, lowest)
				}
			}
		}()
		for {
			msg = <- comms
			if msg[0] == "stop" {
				break
			}
			// area text generator
			area.Update(areaUpdateStr)
			time.Sleep(time.Second / 2)
		}
	}
}
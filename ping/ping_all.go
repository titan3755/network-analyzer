package ping

import (
	"fmt"
	"netzer/utils"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

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
	return nil
}
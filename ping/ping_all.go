package ping

import (
	"fmt"
	"log"
	"net"
	"netzer/utils"
	"strings"
	"sync"
	"time"
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
	gor_ip_latency := make(chan string)
	wg := &sync.WaitGroup{}
	for _, ip := range ip_list {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			var latency string = ""
			startTime := time.Now()
			conn, err := net.Dial("tcp", ip)
			if err != nil {
				latency = fmt.Sprintf("Error: %v", err)
				gor_ip_latency <- latency
			}
			go func() {
				time.Sleep(2 * time.Second)
				if latency == "" || !strings.Contains(latency, "Error") {
					latency = "Error: Timeout"
					gor_ip_latency <- latency
					wg.Done()
				}
			}()
			defer conn.Close()
			latency = fmt.Sprintf("%v", time.Since(startTime))
			log.Default().Printf("Latency: %s, IP: %s\n", latency, ip)
			gor_ip_latency <- latency
		}(ip)
	}
	go func() {
		wg.Wait()
		close(gor_ip_latency)
	}()
	for latency := range gor_ip_latency {
		if strings.Contains(latency, "Error") {
			pterm.Error.Println(fmt.Sprintf("Error: %s", latency))
		} else {
			pterm.Success.Println(fmt.Sprintf("Latency: %s", latency))
		}
	}
	return nil
}

// func pingAll() {
	
// }
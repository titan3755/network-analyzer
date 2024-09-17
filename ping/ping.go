package ping

import (
	"fmt"
	"net"
	"netzer/utils"
	"strings"
	"time"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func PingMain(c *cli.Context) error {
	utils.PingIntro()
	pterm.Info.Println(fmt.Sprintf("Started pinging %v: Press Ctrl+C to stop", c.Args().Get(0)))
	fmt.Println()
	var count_loop int = 0
	var sum_lantency float64 = 0
	var average_latency float64 = 0
	var latency_list []float64
	var highest_latency float64
	var lowest_latency float64
	area, _ := pterm.DefaultArea.Start()
	for {
		latency := PingSpecificIP(c.Args().Get(0))
		if strings.Contains(latency, "Error") {
			area.Update(pterm.Error.Sprintf("Error: %s", latency))
		} else {
			count_loop++
			latency_float, _ := time.ParseDuration(latency)
			sum_lantency += float64(latency_float.Milliseconds())
			average_latency = sum_lantency / float64(count_loop)
			latency_list = append(latency_list, float64(latency_float.Milliseconds()))
			highest_latency = utils.FindMax(latency_list)
			lowest_latency = utils.FindMin(latency_list)
			area.Update(pterm.Success.Sprintf("Latency: %s\nAverage: %s\nHighest: %dms\nLowest: %dms", latency, time.Duration(average_latency) * time.Millisecond, int(highest_latency), int(lowest_latency)))
		}
		time.Sleep(time.Second / 2)
	}
	// return nil
}

func PingSpecificIP(ip string) string {
	var latency string
	startTime := time.Now()
    conn, err := net.Dial("tcp", ip)
    if err != nil {
		latency = fmt.Sprintf("Error: %v", err)
		return latency
    }
    defer conn.Close()
	latency = fmt.Sprintf("%v", time.Since(startTime))
	return latency
}
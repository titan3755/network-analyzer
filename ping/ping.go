package ping

import (
	"fmt"
	"netzer/utils"
	"strings"
	"time"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

// this function pings a specified IP address (main_cmd_function)

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
	var pinger_run bool = true
	area, _ := pterm.DefaultArea.Start()
	ctrl_intr := make(chan os.Signal, 1)
	signal.Notify(ctrl_intr, os.Interrupt)
	go func() {
		for range ctrl_intr {
			pinger_run = false
			area.Clear()
			pterm.Info.Printf("\nTotal pings sent: %d\nAverage latency: %s\nHighest latency: %dms\nLowest latency: %dms\nExiting pinger ...", count_loop, time.Duration(average_latency) * time.Millisecond, int(highest_latency), int(lowest_latency))
		}
	}()
	for pinger_run {
		latency, errt := utils.ICMP_Ping(c.Args().Get(0))
		if errt != nil {
			area.Update(pterm.Error.Sprintf("Error: %v", errt))
		} else if strings.Contains(latency, "Error") {
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
	return nil
}
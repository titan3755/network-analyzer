package ping

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"os"
	"os/signal"
	"strings"
	"time"
)

// this function pings a specified IP address (main_cmd_function)

func PingMain(c *cli.Context) error {
	utils.PingIntro()
	pterm.Info.Println(fmt.Sprintf("Started pinging %v: Press Ctrl+C to stop", c.Args().Get(0)))
	fmt.Println()
	var countLoop = 0
	var sumLantency float64 = 0
	var averageLatency float64 = 0
	var latencyList []float64
	var highestLatency float64
	var lowestLatency float64
	var pingerRun = true
	area, _ := pterm.DefaultArea.Start()
	ctrlIntr := make(chan os.Signal, 1)
	signal.Notify(ctrlIntr, os.Interrupt) // listen for interrupt signal (ctrl+c)
	go func() {
		for range ctrlIntr {
			pingerRun = false
			area.Clear()
			pterm.Info.Printf("\nTotal pings sent: %d\nAverage latency: %s\nHighest latency: %dms\nLowest latency: %dms\nExiting pinger ...", countLoop, time.Duration(averageLatency)*time.Millisecond, int(highestLatency), int(lowestLatency))
		}
	}()
	for pingerRun {
		latency, errt := utils.IcmpPing(c.Args().Get(0))
		if errt != nil {
			area.Update(pterm.Error.Sprintf("Error: %v", errt))
		} else if strings.Contains(latency, "Error") {
			area.Update(pterm.Error.Sprintf("Error: %s", latency))
		} else {
			countLoop++
			latencyFloat, _ := time.ParseDuration(latency)
			sumLantency += float64(latencyFloat.Milliseconds())
			averageLatency = sumLantency / float64(countLoop)
			latencyList = append(latencyList, float64(latencyFloat.Milliseconds()))
			highestLatency = utils.FindMax(latencyList)
			lowestLatency = utils.FindMin(latencyList)
			area.Update(pterm.Success.Sprintf("Latency: %s\nAverage: %s\nHighest: %dms\nLowest: %dms", latency, time.Duration(averageLatency)*time.Millisecond, int(highestLatency), int(lowestLatency)))
		}
		time.Sleep(time.Second / 2)
	}
	return nil
}

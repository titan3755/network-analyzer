package ping

import (
	"fmt"
	"net"
	"netzer/utils"
	"strings"
	"time"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/urfave/cli/v2"
)

func PingIntro() {
	utils.ResetTerminal()
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("NetZer  ", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Ping", pterm.FgLightMagenta.ToStyle()),
	).Render()
	pterm.Info.Println("Welcome to NetZer Ping!")
	pterm.Info.Println("This tool allows you to ping all servers in the IP list or a specific server.")
	fmt.Println()
}

func PingMain(c *cli.Context) error {
	PingIntro()
	if c.Command.Name == "ping-all" || c.Command.Name == "pa" {
		PingAllIP()
	} else if c.Command.Name == "ping-specific-ip" || c.Command.Name == "psi" {
		pterm.Info.Println(fmt.Sprintf("Started pinging %v: Press Ctrl+C to stop", c.Args().Get(0)))
		area, _ := pterm.DefaultArea.Start()
		for {
			latency := PingSpecificIP(c.Args().Get(0))
			if strings.Contains(latency, "Error") {
				area.Update(pterm.Error.Sprintf("Error: %s", latency))
			} else {
				area.Update(pterm.Success.Sprintf("Latency: %s", latency))
			}
			// Pause for 0.5 s before the next update.
			time.Sleep(time.Second / 2)
		}
		area.Stop()
		// end pterm area
	} else if c.Command.Name == "ping" || c.Command.Name == "p" {
		PingAllIP()
	} else {
		pterm.Error.Println("Invalid command. Please try again.")
	}
	return nil
}

func PingAllIP() {
	fmt.Println("Pinging all servers...")
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
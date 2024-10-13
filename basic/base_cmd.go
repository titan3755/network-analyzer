package basic

import (
	"fmt"
	"netzer/utils"

	pterm "github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func BaseCmdMain(c *cli.Context) error {
	utils.BasicIntro()
	pterm.Info.Println("Welcome to Network Analyzer (NetZer) CLI")
	pterm.Info.Println("NetZer is a professional-grade command line interface tool for complete network stability analysis")
	fmt.Print("\n")
	style := pterm.NewStyle(pterm.BgGreen, pterm.FgBlack, pterm.Bold)
	style.Println("Basic Commands -->")
	pterm.Info.Println("Type 'netzer help' to see the list of available commands")
	pterm.Info.Println("Type 'netzer version' to see the version of the application")
	fmt.Print("\n")
	style.Println("Analyzers -->")
	pterm.Info.Println("Type 'netzer stability-analyzer' to run the stability analyzer")
	pterm.Info.Println("Type 'netzer speed-analyzer' to run the speed analyzer")
	pterm.Info.Println("Type 'netzer full-analyzer' to run the full analyzer")
	fmt.Print("\n")
	style.Println("Settings -->")
	pterm.Info.Println("Type 'netzer set-settings' to set the settings")
	pterm.Info.Println("Type 'netzer show-settings' to show the settings")
	pterm.Info.Println("Type 'netzer wipe-settings' to reset the settings")
	fmt.Print("\n")
	style.Println("Ping -->")
	pterm.Info.Println("Type 'netzer ping' to ping a server")
	pterm.Info.Println("Type 'netzer ping-all' to ping all servers in IP file")
	fmt.Print("\n")
	style.Println("IP -->")
	pterm.Info.Println("Type 'netzer add-ip' to add an IP to the IP file")
	// to add more
	return nil
}
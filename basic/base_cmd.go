package basic

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
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
	pterm.Info.Println("Type 'netzer ping-specific-ip' to ping a server")
	pterm.Info.Println("Type 'netzer ping-all-ip' to ping all IPs and hosts in IP file")
	fmt.Print("\n")
	style.Println("IP -->")
	pterm.Info.Println("Type 'netzer add-single-ip-to-file' to add an IP to the IP file")
	pterm.Info.Println("Type 'netzer add-multiple-ip-to-file' to add multiple IPs to the IP file")
	pterm.Info.Println("Type 'netzer remove-single-ip-from-file' to remove an IP from the IP file")
	pterm.Info.Println("Type 'netzer remove-multiple-ip-from-file' to remove multiple IPs from the IP file")
	pterm.Info.Println("Type 'netzer generate-ip-file' to generate an IP file")
	pterm.Info.Println("Type 'netzer use-ip-file' to use an IP file")
	pterm.Info.Println("Type 'netzer read-ip-file' to read an IP file")
	fmt.Print("\n")
	style.Println("Speed Test -->")
	pterm.Info.Println("Type 'netzer speed-test-local' to run a speed test with respect to closest servers")
	pterm.Info.Println("Type 'netzer speed-test-global' to run a speed test with respect to global servers")
	fmt.Print("\n")
	pterm.Info.Println("Enjoy using NetZer!")
	return nil
}

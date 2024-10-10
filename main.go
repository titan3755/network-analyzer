package main

import (
	"log"
	"netzer/analyzers"
	"netzer/basic"
	"netzer/ip"
	"netzer/ping"
	"netzer/utils"
	"netzer/speedtest"
	"os"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

// preChecks function checks if the settings file exists and if not, creates it
// it also checks other pre-requisites

func preChecks() {
	err := utils.SettingsFile()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, ern := utils.GetSettings("ip_file")
	if ern != nil {
		// clear the file
		err := os.WriteFile("settings.prp", []byte(""), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// mainApp function is the main function for the application

func mainApp() {
	app := &cli.App{
		Name: "netzer",
		Usage: "[NETWORK ANALYZER] A network utility tool for analyzing network reliability and stability", 
		Action: func(c *cli.Context) error {
			println("Hello friend!")
			return nil
		},
		Commands: []*cli.Command{
			// {
			// 	Name: "ping",
			// 	Aliases: []string{"p"},
			// 	Usage: "netzer ping utility",
			// 	Action: ping.PingMain,
			// },
			{
				Name: "ping-all-ip",
				Aliases: []string{"pai"},
				Usage: "ping all servers in the IP list",
				Action: ping.PingAllMain,
			},
			{
				Name: "ping-specific-ip",
				Aliases: []string{"psi"},
				Usage: "ping a specific IP address/server address",
				Action: ping.PingMain,
			},
			// {
			// 	Name: "ip",
			// 	Aliases: []string{"i"},
			// 	Usage: "netzer ip utility",
			// 	Action: ip.,
			// },
			{
				Name: "add-single-ip-to-file",
				Aliases: []string{"asif"},
				Usage: "add an IP address to the file",
				Action: ip.AddSingleIPToFileMain,
			},
			{
				Name: "add-multiple-ip-to-file",
				Aliases: []string{"amif"},
				Usage: "add many IP addresses to the file",
				Action: ip.AddMultipleIPToFileMain,
			},
			{
				Name: "remove-single-ip-from-file",
				Aliases: []string{"rsif"},
				Usage: "remove an IP address from the file",
				Action: ip.RemoveSingleIPFromFileMain,
			},
			{
				Name: "remove-multiple-ip-from-file",
				Aliases: []string{"rmif"},
				Usage: "remove multiple IP addresses from the file",
				Action: ip.RemoveMultipleIPFromFileMain,
			},
			{
				Name: "generate-ip-file",
				Aliases: []string{"gif"},
				Usage: "generate a file with a list of IP addresses",
				Action: ip.IPFileGeneratorMain,
			},
			{
				Name: "use-ip-file",
				Aliases: []string{"uif"},
				Usage: "use a IP file at a specified location",
				Action: ip.UseIPFileMain,
			},
			{
				Name: "read-ip-file",
				Aliases: []string{"rdf"},
				Usage: "read the IP file and display the IP addresses",
				Action: ip.ReadIPFromFileMain,
			},
			// {
			// 	Name: "analyzer",
			// 	Aliases: []string{"a"},
			// 	Usage: "netzer analyzer utility",
			// 	Action: analyzers.AnalyzerMain,
			// },
			{
				Name: "stability-analysis",
				Aliases: []string{"sa"},
				Usage: "perform a short network analysis",
				Action: analyzers.StabilityAnalyzerMain,
			},
			{
				Name: "full-analysis",
				Aliases: []string{"fa"},
				Usage: "perform a full network analysis",
				Action: analyzers.StabilityAnalyzerFullMain,
			},
			{
				Name: "long-term-analysis",
				Aliases: []string{"la"},
				Usage: "perform a lengthy network analysis",
				Action: analyzers.StabilityAnalyzerLongMain,
			},
			{
				Name: "speed-test-local",
				Aliases: []string{"stl"},
				Usage: "perform a speed test with respect to a domestic/local server",
				Action: speedtest.SpeedTestLocalMain,
			},
			{
				Name: "speed-test-global",
				Aliases: []string{"stg"},
				Usage: "perform a speed test with respect to any server in the world",
				Action: speedtest.SpeedTestGlobalMain,
			},
			{
				Name: "help",
				Aliases: []string{"h"},
				Usage: "show help",
				Action: basic.ShowHelp,
			},
			{
				Name: "version",
				Aliases: []string{"v"},
				Usage: "show cli version",
				Action: basic.ShowVersion,
			},
			{
				Name: "show-settings",
				Aliases: []string{"shs"},
				Usage: "show settings",
				Action: basic.ShowSettingsMain,
			},
			{
				Name: "set-settings",
				Aliases: []string{"sts"},
				Usage: "set settings",
				Action: basic.SetSettingsMain,
			},
			{
				Name: "wipe-settings",
				Aliases: []string{"sws"},
				Usage: "wipe settings",
				Action: basic.WipeSettingsMain,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// main function

func main()  {
	preChecks()
	mainApp()
}
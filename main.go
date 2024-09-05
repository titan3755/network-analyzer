package main

import (
	"netzer/ping"
	"netzer/utils"
	"netzer/ip"
	"log"
	"os"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func main()  {
	err := utils.SettingsFile()
	if err != nil {
		log.Fatal(err)
		return
	}
	app := &cli.App{
		Name: "netzer",
		Usage: "[NETWORK ANALYZER] A network utility tool for analyzing network reliability and stability", 
		Action: func(c *cli.Context) error {
			println("Hello friend!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "ping",
				Aliases: []string{"p"},
				Usage: "netzer ping utility",
				Action: ping.PingMain,
			},
			{
				Name: "ping-all",
				Aliases: []string{"pa"},
				Usage: "ping all servers in the IP list",
				Action: ping.PingMain,
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
			{
				Name: "stability-analysis",
				Aliases: []string{"sa"},
				Usage: "analyze network stability",
				Action: analyzeNetworkStability,
			},
			{
				Name: "full-analysis",
				Aliases: []string{"fa"},
				Usage: "perform a full network analysis",
				Action: fullNetworkAnalysis,
			},
			{
				Name: "long-term-analysis",
				Aliases: []string{"la"},
				Usage: "perform a long term network analysis",
				Action: analyzeNetworkStabilityLT,
			},
			{
				Name: "speed-test",
				Aliases: []string{"st"},
				Usage: "perform a speed test",
				Action: speedTest,
			},
			{
				Name: "help",
				Aliases: []string{"h"},
				Usage: "show help",
				Action: showHelp,
			},
			{
				Name: "version",
				Aliases: []string{"v"},
				Usage: "show cli version",
				Action: showVersion,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func analyzeNetworkStability(c *cli.Context) error {
	println("Analyzing network stability")
	return nil
}

func fullNetworkAnalysis(c *cli.Context) error {
	println("Performing a full network analysis")
	return nil
}

func speedTest(c *cli.Context) error {
	println("Performing a speed test")
	return nil
}

func analyzeNetworkStabilityLT(c *cli.Context) error {
	println("Performing a long term network analysis")
	return nil
}

func showHelp(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}

func showVersion(c *cli.Context) error {
	println("Version 1.0.0")
	return nil
}
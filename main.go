package main

import (
	"os"
	"log"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func main()  {
	app := &cli.App{
		Name: "netzer",
		Usage: "[NETWORK ANALYZER] A network utility tool for analyzing network reliability and stability", 
		Action: func(c *cli.Context) error {
			println("Hello friend!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "ping-all",
				Aliases: []string{"pa"},
				Usage: "ping all servers in the IP list",
				Action: pingAllServers,
			},
			{
				Name: "stability-analyzer",
				Aliases: []string{"sa"},
				Usage: "analyze network stability",
				Action: analyzeNetworkStability,
			},
			{
				Name: "speed-test",
				Aliases: []string{"st"},
				Usage: "perform a speed test",
				Action: speedTest,
			},
			{
				Name: "full-analysis",
				Aliases: []string{"fa"},
				Usage: "perform a full network analysis",
				Action: analyzeNetworkStability,
			},
			{
				Name: "long-term-analysis",
				Aliases: []string{"lta"},
				Usage: "perform a long term network analysis",
				Action: analyzeNetworkStabilityLT,
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
			{
				Name: "ping-specific-ip",
				Aliases: []string{"psi"},
				Usage: "ping a specific IP address/server address",
				Action: pingSpecificIP,
			},
			{
				Name: "add-ip-to-list",
				Aliases: []string{"ail"},
				Usage: "add an IP address to the list",
				Action: addIP,
			},
			{
				Name: "remove-ip-from-list",
				Aliases: []string{"ril"},
				Usage: "remove an IP address from the list",
				Action: removeIP,
			},
			{
				Name: "generate-ip-file",
				Aliases: []string{"gif"},
				Usage: "generate a file with a list of IP addresses",
				Action: genIPFile,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func pingAllServers(c *cli.Context) error {
	println("Pinging all servers in the IP list")
	return nil
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

func pingSpecificIP(c *cli.Context) error {
	println("Pinging a specific IP address/server address")
	return nil
}

func addIP(c *cli.Context) error {
	println("Adding an IP address to the list")
	return nil
}

func removeIP(c *cli.Context) error {
	println("Removing an IP address from the list")
	return nil
}

func genIPFile(c *cli.Context) error {
	println("Generating a file with a list of IP addresses")
	return nil
}
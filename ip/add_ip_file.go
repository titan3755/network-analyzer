package ip

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
	"github.com/pterm/pterm"
	"netzer/utils"
)

func AddIPToFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println(fmt.Sprintf("Adding IP address %v to the list...", c.Args().First()))
	if c.Args().First() == "" {
		var error_txt string = "error: no ip address provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	} else if c.Args().Get(1) == "" {
		var error_txt string = "error: no file path provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	}
	err := addIPToFile(c.Args().First(), c.Args().Get(1))
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP address added successfully!")
	return nil
}

func addIPToFile(ip string, location string) error {
	// check if IP address is valid
	if !utils.CheckIfValidIPv4(ip) {
		return fmt.Errorf("invalid ip address")
	}
	// check if file exists
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	// add IP address to file
	file, err := os.OpenFile(location, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(ip + ";")
	if err != nil {
		return err
	}
	return nil
}
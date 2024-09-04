package ip

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
	"github.com/pterm/pterm"
	"netzer/utils"
)

func AddSingleIPToFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println(fmt.Sprintf("Adding IP address %v to the list...", c.Args().Get(1)))
	if c.Args().Get(1) == "" {
		var error_txt string = "error: no ip address provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	} else if c.Args().First() == "" {
		var error_txt string = "error: no file path provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	}
	err := addIPToFile(c.Args().Get(1), c.Args().First())
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP address added successfully!")
	return nil
}

func AddMultipleIPToFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println("Adding IP addresses to the list...")
	if c.Args().Get(1) == "" {
		var error_txt string = "error: no ip addresses provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	} else if c.Args().First() == "" {
		var error_txt string = "error: no file path provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	}
	var ipList []string
	for _, ip := range c.Args().Slice() {
		ern := utils.CheckIfValidIPv4(ip)
		if ern {
			ipList = append(ipList, ip)	
		} else {
			pterm.Error.Println(fmt.Sprintf("invalid IP address %v provided, not added to file", ip))
		}
	}
	err := addMultipleIPToFile(utils.ConvListOfIPToFileFormat(ipList), c.Args().First())
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP addresses added successfully!")
	return nil
}

func addMultipleIPToFile(ipConvFormat string, location string) error {
	// check if file exists
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	// check if IP addresses are already in file
	if in, err := utils.CheckIfIPAlreadyInFile(ipConvFormat, location); err != nil {
		return err
	} else if in {
		return fmt.Errorf("ip addresses already in file")
	}
	// add IP addresses to file
	file, err := os.OpenFile(location,
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(ipConvFormat)
	if err != nil {
		return err
	}
	return nil
}


func addIPToFile(ip string, location string) error {
	// check if IP address is valid
	if !utils.CheckIfValidIPv4(ip) {
		return fmt.Errorf("invalid ip address")
	}
	// check if IP address is already in file
	if in, err := utils.CheckIfIPAlreadyInFile(ip, location); err != nil {
		return err
	} else if in {
		return fmt.Errorf("ip address already in file")
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
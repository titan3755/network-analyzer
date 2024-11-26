package ip

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"os"
)

// this function adds a single IP address to a file (main_cmd_function)

func AddSingleIPToFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println(fmt.Sprintf("Adding IP address %v to the list...", c.Args().Get(1)))
	if c.Args().Get(1) == "" {
		var errorTxt = "error: no ip address provided"
		pterm.Error.Println(errorTxt)
		return fmt.Errorf("%s", errorTxt)
	} else if c.Args().First() == "" {
		var errorTxt = "error: no file path provided"
		pterm.Error.Println(errorTxt)
		return fmt.Errorf("%s", errorTxt)
	}
	err := addIPToFile(c.Args().Get(1), c.Args().First())
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP address added successfully!")
	return nil
}

// this function adds multiple IP addresses to a file (main_cmd_function)

func AddMultipleIPToFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println("Adding IP addresses to the list...")
	if c.Args().Get(1) == "" {
		var errorTxt = "error: no ip addresses provided"
		pterm.Error.Println(errorTxt)
		return fmt.Errorf("%s", errorTxt)
	} else if c.Args().First() == "" {
		var errorTxt = "error: no file path provided"
		pterm.Error.Println(errorTxt)
		return fmt.Errorf("%s", errorTxt)
	}
	var ipList []string
	for _, ip := range c.Args().Slice() {
		if ip == c.Args().First() {
			continue
		}
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

// this function adds multiple IP addresses to a file

func addMultipleIPToFile(ipConvFormat string, location string) error {
	// check if file format is valid if file is not empty
	if dat, ern := os.ReadFile(location); ern != nil {
		return ern
	} else if string(dat) != "" {
		if err := utils.CheckIPFileFormatValidity(location); err != nil {
			return err
		}
	}
	// check if file exists
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	// check if IP addresses are already in file
	ipNlist := utils.ConvFileFormatToListOfIP(ipConvFormat)
	for _, ip := range ipNlist {
		if in, err := utils.CheckIfIPAlreadyInFile(ip, location); err != nil {
			return err
		} else if in && ip != "" {
			return fmt.Errorf("ip address(es) already in file")
		}
	}
	// add IP addresses to file
	file, err := os.OpenFile(location,
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	_, err = file.WriteString(ipConvFormat)
	if err != nil {
		return err
	}
	return nil
}

// this function adds a single IP address to a file

func addIPToFile(ip string, location string) error {
	// check if file format is valid if file is not empty
	if dat, ern := os.ReadFile(location); ern != nil {
		return ern
	} else if string(dat) != "" {
		if err := utils.CheckIPFileFormatValidity(location); err != nil {
			return err
		}
	}
	// check if IP address is valid
	if !utils.CheckIfValidIPv4(ip) {
		return fmt.Errorf("invalid ip address")
	}
	// check if IP address is already in file
	if in, err := utils.CheckIfIPAlreadyInFile(ip, location); err != nil {
		return err
	} else if in {
		return fmt.Errorf("ip address(es) already in file")
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	_, err = file.WriteString(ip + ";")
	if err != nil {
		return err
	}
	return nil
}

package ip

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
	"github.com/pterm/pterm"
	"netzer/utils"
	"strings"
)

func RemoveSingleIPFromFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println(fmt.Sprintf("Removing IP address %v from the list...", c.Args().Get(1)))
	if c.Args().Get(1) == "" {
		var error_txt string = "error: no ip address provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	} else if c.Args().First() == "" {
		var error_txt string = "error: no file path provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	}
	err := removeIPFromFile(c.Args().Get(1), c.Args().First())
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP address removed successfully!")
	return nil
}

func RemoveMultipleIPFromFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println("Removing IP addresses from the list...")
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
		if ip == c.Args().First() || ip == "" {
			continue
		}
		ern := utils.CheckIfValidIPv4(ip)
		if ern {
			ipList = append(ipList, ip)	
		} else {
			pterm.Error.Println(fmt.Sprintf("invalid IP address %v provided, not removed from file", ip))
		}
	}
	if len(ipList) == 0 {
		var error_txt string = "error: no valid ip addresses provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	}
	err := removeMultipleIPFromFile(utils.ConvListOfIPToFileFormat(ipList), c.Args().First())
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP addresses removed successfully!")
	return nil
}

func removeIPFromFile(ip string, fileLocation string) error {
	// check if file exists
	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	// check if file is in correct format
	errn := utils.CheckIPFileFormatValidity(fileLocation)
	if errn != nil {
		return errn
	}
	// open and read file at location
	chk, err := utils.CheckIfIPAlreadyInFile(ip, fileLocation)
	if err != nil {
		return err
	}
	if chk {
		// open and read file at location
		dat, err := os.ReadFile(fileLocation)
		if err != nil {
			return err
		}
		// remove the ip from data
		if string(dat) == "" {
			return nil
		}
		var newDat string
		for _, line := range strings.Split(string(dat), ";") {
			if strings.Contains(line, ip) {
				continue
			}
			newDat += line + ";"
		}
		file, er := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if er != nil {
			return er
		}
		defer file.Close()
		_, err = file.WriteString(newDat)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("ip address not in file")
}

func removeMultipleIPFromFile(ipConvFormat string, location string) error {
	// check if file exists
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	// check if file is in correct format
	errn := utils.CheckIPFileFormatValidity(location)
	if errn != nil {
		return errn
	}
	// open and read file at location
	dat, err := os.ReadFile(location)
	if err != nil {
		return err
	}
	// remove the ip from data
	if string(dat) == "" {
		return nil
	}
	var newDat string
	for _, line := range strings.Split(string(dat), ";") {
		if strings.Contains(ipConvFormat, line) || line == "" {
			continue
		}
		newDat += line + ";"
	}
	file, er := os.OpenFile(location,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if er != nil {
		return er
	}
	defer file.Close()
	_, err = file.WriteString(newDat)
	if err != nil {
		return err
	}
	return nil
}
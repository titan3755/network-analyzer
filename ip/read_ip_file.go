package ip

import (
	"fmt"
	"netzer/utils"
	"os"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

// this function reads IP addresses from a file after checking its validity (main_cmd_function)
func ReadIPFromFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println(fmt.Sprintf("Reading IP addresses from file at %v ...", c.Args().First()))
	if c.Args().Get(0) == "" {
		var error_txt string = "error: no file path provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	}
	ip_lst, err := readIPFromFile(c.Args().First())
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return fmt.Errorf("%v", err)
	}
	pterm.Success.Println("IP addresses read successfully!")
	pterm.Info.Println("IP addresses:")
	for _, ip := range ip_lst {
		if ip == "" {
			continue
		} else {
			pterm.Info.Println(ip)
		}
	}
	return nil
}

// this function reads IP addresses from a file after checking its validity

func readIPFromFile(filePath string) ([]string, error) {
	// check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []string{}, fmt.Errorf("file does not exist")
	}
	// check if file is in correct format
	errn := utils.CheckIPFileFormatValidity(filePath)
	if errn != nil {
		return []string{}, errn
	}
	// check and remove duplicate ip addresses
	err := utils.RemoveDuplicateIPFromFile(filePath)
	if err != nil {
		return []string{}, err
	}
	// read file
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return []string{}, err
	}
	// parse ip addresses
	ipList := utils.ConvFileFormatToListOfIP(string(dat))
	return ipList, nil
}
package ip

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"os"
)

// this function sets the ip_file setting in the settings.prp file to the specified directory in order to
// use it for analyzing cmds and ping cmds (main_cmd_function)

func UseIPFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println("Using IP file...")
	if c.Args().First() == "" {
		var errorTxt = "error: no file path provided"
		pterm.Error.Println(errorTxt)
		return fmt.Errorf("%s", errorTxt)
	}
	err := useIPFile(c.Args().First())
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP file used successfully!")
	return nil
}

// this function sets the ip_file setting in the settings.prp file to the specified directory in order to
// use it for analyzing cmds and ping cmds

func useIPFile(filePath string) error {
	// check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	// check if file is in correct format
	errn := utils.CheckIPFileFormatValidity(filePath)
	if errn != nil {
		return errn
	}
	// check and remove duplicate ip addresses
	errns := utils.RemoveDuplicateIPFromFile(filePath)
	if errns != nil {
		return errns
	}
	// set file path as currently used ip file in the settings.pfp
	err := utils.SetSettings("ip_file", filePath)
	if err != nil {
		return err
	}
	pterm.Success.Println(fmt.Sprintf("IP file %v is now being used", filePath))
	return nil
}

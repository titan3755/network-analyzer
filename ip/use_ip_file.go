package ip

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
	"github.com/pterm/pterm"
	"netzer/utils"
)

func UseIPFileMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println("Using IP file...")
	if c.Args().First() == "" {
		var error_txt string = "error: no file path provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	}
	err := useIPFile(c.Args().First())
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP file used successfully!")
	return nil
}

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
	// set file path as currently used ip file in the settings.pfp
	err := utils.SetSettings("ip_file", filePath)
	if err != nil {
		return err
	}
	pterm.Success.Println(fmt.Sprintf("IP file %v is now being used", filePath))
	return nil
}
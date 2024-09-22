package ip

import (
	"os"
	"netzer/utils"
	"github.com/urfave/cli/v2"
	"github.com/pterm/pterm"
	"fmt"
)

// this function generates an IP file in the specified directory (main_cmd_function)

func IPFileGeneratorMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println(fmt.Sprintf("Generating IP file at %v", c.Args().First()))
	if c.Args().First() == "" {
		var error_txt string = "error: no path provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	} else if c.Args().Get(1) == "" {
		var error_txt string = "error: no file name provided"
		pterm.Error.Println(error_txt)
		return fmt.Errorf("%s", error_txt)
	}
	err := ipFileGenerator(c.Args().First(), c.Args().Get(1))
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		return err
	}
	pterm.Success.Println("IP file generated successfully!")
	return nil
}

// this function generates an IP file in the specified directory

func ipFileGenerator(path string, fileName string) error {
	file, err := os.Create(path + "/" + fileName + ".ip")
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
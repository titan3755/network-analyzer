package ip

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
	"os"
)

// this function generates an IP file in the specified directory (main_cmd_function)

func IPFileGeneratorMain(c *cli.Context) error {
	utils.IPIntro()
	pterm.Info.Println(fmt.Sprintf("Generating IP file at %v", c.Args().First()))
	if c.Args().First() == "" {
		var errorTxt = "error: no path provided"
		pterm.Error.Println(errorTxt)
		return fmt.Errorf("%s", errorTxt)
	} else if c.Args().Get(1) == "" {
		var errorTxt = "error: no file name provided"
		pterm.Error.Println(errorTxt)
		return fmt.Errorf("%s", errorTxt)
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			pterm.Error.Println(fmt.Sprintf("Error: %v", err))
		}
	}(file)
	return nil
}

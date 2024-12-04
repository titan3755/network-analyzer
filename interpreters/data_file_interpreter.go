package interpreters

import (
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"netzer/utils"
)

func DataFileInterpreterMain(c *cli.Context) error {
	utils.InterpreterIntro()
	var mode string
	var filePath string
	if c.Args().Get(0) == "" {
		filePath = "stability_test_data.data"
		pterm.Info.Println("No file path provided, using default file path [./stability_test_data.data]...")
	} else {
		filePath = c.Args().Get(0)
		pterm.Info.Println("Using file path: ", c.Args().Get(0))
	}
	if c.Args().Get(1) == "" {
		mode = "stb"
		pterm.Info.Println("No mode provided, using default mode [stb -> stability test file]...")
	} else {
		mode = c.Args().Get(1)
		pterm.Info.Println("Using mode: ", c.Args().Get(1))
	}
	// check if file exists or not
	// also check if file empty
	// if file empty, return error
	isEmpty, err := utils.FileEmptyCheck(filePath)
	if err != nil {
		pterm.Error.Println("Error: ", err)
		return err
	}
	if isEmpty {
		pterm.Error.Println("Error: file is empty")
		return nil
	}
	// check if mode is valid
	if mode != "stb" && mode != "spd" {
		pterm.Error.Println("Error: invalid mode provided")
		return nil
	}
	// read file
	pterm.Info.Println("Checks completed successfully ...")
	pterm.Info.Println("Reading file ...")
	if mode == "stb" {
		data := utils.ReadAnalyzerStabilityTestDataFromFile(filePath)
		utils.ShowAnalyzerStabilityTestData(data)
	} else {
		data := utils.ReadAnalyzerSpeedTestDataFromFile(filePath)
		utils.ShowAnalyzerSpeedTestData(data)
	}
	return nil
}

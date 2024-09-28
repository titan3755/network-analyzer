package analyzers

import (
	"github.com/urfave/cli/v2"
	"netzer/utils"
)

func StabilityAnalyzerLongMain(c *cli.Context) error {
	utils.AnalyzerIntro()
	return nil
}
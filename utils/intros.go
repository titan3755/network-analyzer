package utils

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"fmt"
)

func PingIntro() {
	ResetTerminal()
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("NetZer  ", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Ping", pterm.FgLightMagenta.ToStyle()),
	).Render()
	pterm.Info.Println("Welcome to NetZer Ping!")
	pterm.Info.Println("This utility allows you to ping all servers in the IP list or a specific server.")
	fmt.Println()
}

func IPIntro() {
	ResetTerminal()
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("NetZer  ", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("IP", pterm.FgLightMagenta.ToStyle()),
	).Render()
	pterm.Info.Println("Welcome to NetZer IP!")
	pterm.Info.Println("This utility allows you to generate new IP address file, modify it, use a IP file at a specified location and add or remove IP addresses from the file.")
	fmt.Println()
}

func AnalyzerIntro() {
	ResetTerminal()
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("NetZer  ", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Analyzer", pterm.FgLightMagenta.ToStyle()),
	).Render()
	pterm.Info.Println("Welcome to NetZer Analyzer!")
	pterm.Info.Println("This utility allows you to analyze network reliability and stability.")
	fmt.Println()
}

func BasicIntro() {
	ResetTerminal()
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("NetZer  ", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Basic", pterm.FgLightMagenta.ToStyle()),
	).Render()
	pterm.Info.Println("Welcome to NetZer Basic!")
	pterm.Info.Println("This utility allows you to view and modify basic settings and execute other simple commands.")
	fmt.Println()
}
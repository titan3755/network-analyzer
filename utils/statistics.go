package utils

import (
	"fmt"
	"os"
	"strconv"
	"netzer/data"
	tb "github.com/aquasecurity/table"
	pterm "github.com/pterm/pterm"
)

func StatisticsTableCreatorForPingAll(ip_map map[string][][]string, error_map map[string][]error) {
	DBStatisticsTableIntro()
	tableIP := tb.New(os.Stdout)
	tableIP.SetHeaders("IP Address", "Packets Sent", "Packets Received", "Packet Loss", "Min RTT", "Max RTT", "Avg RTT")
	tableIP.SetAlignment(tb.AlignCenter)
	for ip, data := range ip_map {
		latest := data[len(data)-1]
		avgRTT := latest[1]
		maxRTT := latest[2]
		minRTT := latest[3]
		pktSent := latest[4]
		pktRecv := latest[5]
		pktLoss := latest[6]
		tableIP.AddRow(ip, pktSent, pktRecv, pktLoss, minRTT, maxRTT, avgRTT)
	}
	tableIP.Render()
	var noErr bool = true
	tableErr := tb.New(os.Stdout)
	tableErr.SetHeaders("IP Address", "Error")
	tableErr.SetAlignment(tb.AlignCenter)
	for ip, errs := range error_map {
		for _, err := range errs {
			noErr = false
			tableErr.AddRow(ip, fmt.Sprintf("%s", err))
		}
	}
	if noErr {
		tableErr.AddRow("No errors", "No errors")
	}
	tableErr.Render()
}

func StatisticsTableCreatorForStabilityAnalyzer(ip_map map[string][][]string, error_map map[string][]error) {
	DBStatisticsTableIntro()
	var ipIndividualStabilityGradeData map[string]string = make(map[string]string)
	fmt.Print("\n")
	pterm.Info.Println("Ping statistics for the IP addresses:")
	fmt.Print("\n")
	// ping stats
	tableIP := tb.New(os.Stdout)
	tableIP.SetHeaders("IP Address", "Packets Sent", "Packets Received", "Packet Loss", "Min RTT", "Max RTT", "Avg RTT")
	tableIP.SetAlignment(tb.AlignCenter)
	for ip, data := range ip_map {
		latest := data[len(data)-1]
		avgRTT := latest[1]
		maxRTT := latest[2]
		minRTT := latest[3]
		pktSent := latest[4]
		pktRecv := latest[5]
		pktLoss := latest[6]
		tableIP.AddRow(ip, pktSent, pktRecv, pktLoss, minRTT, maxRTT, avgRTT)
	}
	tableIP.Render()
	var noErr bool = true
	fmt.Print("\n")
	pterm.Info.Println("Error messages for the IP addresses:")
	fmt.Print("\n")
	// error stats
	tableErr := tb.New(os.Stdout)
	tableErr.SetHeaders("IP Address", "Error")
	tableErr.SetAlignment(tb.AlignCenter)
	for ip, errs := range error_map {
		for _, err := range errs {
			noErr = false
			tableErr.AddRow(ip, fmt.Sprintf("%s", err))
		}
	}
	if noErr {
		tableErr.AddRow("No errors", "No errors")
	}
	tableErr.Render()
	fmt.Print("\n")
	pterm.Info.Println("Stability statistics for the IP addresses:")
	fmt.Print("\n")
	// stability stats
	tableStab := tb.New(os.Stdout)
	tableStab.SetHeaders("IP Address", "Recv/Sent %", "\u0394Ping", "Avg", "Stability Grade")
	tableStab.SetAlignment(tb.AlignCenter)
	for ip, data := range ip_map {
		latest := data[len(data)-1]
		avgRTT := latest[1]
		maxRTT := latest[2]
		minRTT := latest[3]
		pktSent := latest[4]
		pktRecv := latest[5]
		pktLoss := latest[6]
		grade, er := CalculateStabilityGrade(pktSent, pktRecv, pktLoss, minRTT, maxRTT, avgRTT)
		if er != nil {
			tableStab.AddRow(ip, pktRecv+"/"+pktSent, "N/A", "N/A", "N/A")
			continue
		}
		// remove "ms" from minrtt, maxrtt and avgrtt
		minRTT = minRTT[:len(minRTT)-2]
		maxRTT = maxRTT[:len(maxRTT)-2]
		// convert maxRTT, minRTT, pktRecv, pktSent to int
		maxRTTInt, err := strconv.ParseFloat(maxRTT, 64)

		if err != nil {
			tableStab.AddRow(ip, pktRecv+"/"+pktSent, "N/A", "N/A", "N/A")
			continue
		}
		minRTTInt, err := strconv.ParseFloat(minRTT, 64)
		if err != nil {
			tableStab.AddRow(ip, pktRecv+"/"+pktSent, "N/A", "N/A", "N/A")
			continue
		}
		pktRecvInt, err := strconv.ParseFloat(pktRecv, 64)
		if err != nil {
			tableStab.AddRow(ip, pktRecv+"/"+pktSent, "N/A", "N/A", "N/A")
			continue
		}
		pktSentInt, err := strconv.ParseFloat(pktSent, 64)
		if err != nil {
			tableStab.AddRow(ip, pktRecv+"/"+pktSent, "N/A", "N/A", "N/A")
			continue
		}
		// calculate delta ping
		deltaPing := maxRTTInt - minRTTInt
		// calculate recv/sent %
		recvSent := (pktRecvInt / pktSentInt) * 100
		tableStab.AddRow(ip, fmt.Sprintf("%.2f", recvSent), fmt.Sprintf("%.2f", deltaPing), avgRTT, grade)
		ipIndividualStabilityGradeData[ip] = grade
	}
	tableStab.Render()
	// stability grade
	overallStabilityGrade := CalculateOverallStabilityGrade(ipIndividualStabilityGradeData)
	fmt.Print("\n")
	pterm.Info.Println("The stability grade is calculated according to the following table:")
	fmt.Print("\n")
	// stability grade table
	tableStabGradeDemonstration := tb.New(os.Stdout)
	tableStabGradeDemonstration.SetHeaders("Stability Grade", "Description")
	tableStabGradeDemonstration.SetAlignment(tb.AlignCenter)
	for i := 0; i <= 9; i++ {
		tableStabGradeDemonstration.AddRow(data.StabilityGrade[i], data.StabilityGradeDescription[data.StabilityGrade[i]])
	}
	tableStabGradeDemonstration.Render()
	fmt.Print("\n")
	pterm.Info.Println("Overall Network Stability Grade:", overallStabilityGrade)
	// graph confirmation
	fmt.Print("\n")
	res, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Do you want to see the stability grade graph?").WithConfirmText("Yes").Show()
	if res {
		GenerateStabilityGradeGraph(ipIndividualStabilityGradeData)
	}
}
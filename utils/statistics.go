package utils

import (
	"fmt"
	"os"
	"strconv"
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
	pterm.Info.Println("\nOverall stability grade:", overallStabilityGrade)
}
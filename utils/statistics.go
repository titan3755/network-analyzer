package utils

import (
	"fmt"
	"os"

	tb "github.com/aquasecurity/table"
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
	tableErr := tb.New(os.Stdout)
	tableErr.SetHeaders("IP Address", "Error")
	tableErr.SetAlignment(tb.AlignCenter)
	for ip, errs := range error_map {
		for _, err := range errs {
			tableErr.AddRow(ip, fmt.Sprintf("%s", err))
		}
	}
	tableErr.Render()
}
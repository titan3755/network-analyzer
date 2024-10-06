package utils

import (
	"net"
	pterm "github.com/pterm/pterm"
)

func GenerateStabilityGradeGraph(ip_graph_data map[string]string) {
	var gradeLst []string
	var hst []string
	for ip, grade := range ip_graph_data {
		host, erh := net.LookupAddr(ip)
		if erh != nil {
			continue
		}
		gradeLst = append(gradeLst, grade)
		hst = append(hst, host[0])
	}
	// Create a new line chart
	bar := []pterm.Bar{}
	var count int = 0
	for i, host := range hst {
		br := pterm.Bar{
			Label: host,
			Value: gradeToNumber(gradeLst[i]),
		}
		bar = append(bar, br)
		count++
	}
	if count == 0 {
		pterm.Error.Println("No data to graph")
		return
	}
	// Create a new line chart
	pterm.DefaultBarChart.WithHorizontal().WithBars(bar).WithHeight(4).Render()
}

func gradeToNumber(grade string) int {
	switch grade {
	case "A":
		return 10
	case "B":
		return 9
	case "C":
		return 8
	case "D":
		return 7
	case "F":
		return 6
	case "G":
		return 5
	case "H":
		return 4
	case "I":
		return 3
	case "J":
		return 2
	case "K":
		return 1
	default:
		return 0
	}
}


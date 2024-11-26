package utils

import (
	"math"
	"netzer/data"
	"strconv"
)

func CalculateStabilityGrade(psent, precv, ploss, minrtt, maxrtt, avgrtt string) (string, error) {
	var stabGrade = data.StabilityGrade
	// remove "ms" from minrtt, maxrtt and avgrtt
	minrtt = minrtt[:len(minrtt)-2]
	maxrtt = maxrtt[:len(maxrtt)-2]
	avgrtt = avgrtt[:len(avgrtt)-2]
	// convert string to float64
	// sent, err := strconv.ParseFloat(psent, 64)
	// if err != nil {
	// 	return "", err
	// }
	// recv, err := strconv.ParseFloat(precv, 64)
	// if err != nil {
	// 	return "", err
	// }
	loss, err := strconv.ParseFloat(ploss, 64)
	if err != nil {
		return "", err
	}
	minN, err := strconv.ParseFloat(minrtt, 64)
	if err != nil {
		return "", err
	}
	maxN, err := strconv.ParseFloat(maxrtt, 64)
	if err != nil {
		return "", err
	}
	avg, err := strconv.ParseFloat(avgrtt, 64)
	if err != nil {
		return "", err
	}
	// calculate stability grade
	var lossGrade string
	var rttGrade string
	var deltaPingGrade string
	var grade string
	if loss >= 0.0 {
		if loss < 10 {
			lossGrade = stabGrade[0]
		} else if loss >= 10 && loss < 20 {
			lossGrade = stabGrade[1]
		} else if loss >= 20 && loss < 30 {
			lossGrade = stabGrade[2]
		} else if loss >= 30 && loss < 40 {
			lossGrade = stabGrade[3]
		} else if loss >= 40 && loss < 50 {
			lossGrade = stabGrade[4]
		} else if loss >= 50 && loss < 60 {
			lossGrade = stabGrade[5]
		} else if loss >= 60 && loss < 70 {
			lossGrade = stabGrade[6]
		} else if loss >= 70 && loss < 80 {
			lossGrade = stabGrade[7]
		} else if loss >= 80 && loss < 90 {
			lossGrade = stabGrade[8]
		} else if loss >= 90 && loss <= 100 {
			lossGrade = stabGrade[9]
		} else {
			lossGrade = stabGrade[9]
		}
	}
	if avg >= 0.0 {
		if avg < 50 {
			rttGrade = stabGrade[0]
		} else if avg >= 50 && avg < 100 {
			rttGrade = stabGrade[1]
		} else if avg >= 100 && avg < 150 {
			rttGrade = stabGrade[2]
		} else if avg >= 150 && avg < 200 {
			rttGrade = stabGrade[3]
		} else if avg >= 200 && avg < 250 {
			rttGrade = stabGrade[4]
		} else if avg >= 250 && avg < 300 {
			rttGrade = stabGrade[5]
		} else if avg >= 300 && avg < 350 {
			rttGrade = stabGrade[6]
		} else if avg >= 350 && avg < 400 {
			rttGrade = stabGrade[7]
		} else if avg >= 400 && avg < 450 {
			rttGrade = stabGrade[8]
		} else if avg >= 450 && avg <= 500 {
			rttGrade = stabGrade[9]
		} else {
			rttGrade = stabGrade[9]
		}
	}
	if minN >= 0.0 && maxN >= 0.0 {
		if maxN-minN < 50 {
			deltaPingGrade = stabGrade[0]
		} else if maxN-minN >= 50 && maxN-minN < 100 {
			deltaPingGrade = stabGrade[1]
		} else if maxN-minN >= 100 && maxN-minN < 150 {
			deltaPingGrade = stabGrade[2]
		} else if maxN-minN >= 150 && maxN-minN < 200 {
			deltaPingGrade = stabGrade[3]
		} else if maxN-minN >= 200 && maxN-minN < 250 {
			deltaPingGrade = stabGrade[4]
		} else if maxN-minN >= 250 && maxN-minN < 300 {
			deltaPingGrade = stabGrade[5]
		} else if maxN-minN >= 300 && maxN-minN < 350 {
			deltaPingGrade = stabGrade[6]
		} else if maxN-minN >= 350 && maxN-minN < 400 {
			deltaPingGrade = stabGrade[7]
		} else if maxN-minN >= 400 && maxN-minN < 450 {
			deltaPingGrade = stabGrade[8]
		} else if maxN-minN >= 450 && maxN-minN <= 500 {
			deltaPingGrade = stabGrade[9]
		} else {
			deltaPingGrade = stabGrade[9]
		}
	}
	// calculate average grade from loss, rtt, deltaPing
	total := 0
	for i, s := range stabGrade {
		if s == lossGrade {
			total += i + 1
			break
		}
	}
	for i, s := range stabGrade {
		if s == rttGrade {
			total += i + 1
			break
		}
	}
	for i, s := range stabGrade {
		if s == deltaPingGrade {
			total += i + 1
			break
		}
	}
	indexNo := math.Ceil(float64(total) / 3)
	if indexNo > 10 {
		indexNo = 10
	} else if indexNo < 1 {
		indexNo = 1
	}
	grade = stabGrade[int(indexNo)-1]
	return grade, nil
}

func CalculateOverallStabilityGrade(ipData map[string]string) string {
	var stabGrade = data.StabilityGrade
	// calculate overall stability grade
	var grade string
	var total float64
	var count int
	for _, g := range ipData {
		for i, s := range stabGrade {
			if g == s {
				total += float64(i + 1)
				count++
				break
			}
		}
	}
	if count == 0 {
		return "N/A"
	}
	indexNo := math.Ceil(total / float64(count))
	if indexNo > 10 {
		indexNo = 10
	} else if indexNo < 1 {
		indexNo = 1
	}
	grade = stabGrade[int(indexNo)-1]
	return grade
}

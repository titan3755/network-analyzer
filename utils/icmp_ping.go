package utils

import (
	"fmt"
	probing "github.com/prometheus-community/pro-bing"
	"time"
)

// ICMP_Ping is a function that sends an ICMP echo request to the specified host

func IcmpPing(host string) (latency string, er error) {
	pinger, err := probing.NewPinger(host)
	pinger.SetPrivileged(true)
	if err != nil {
		return "", err
	}
	pinger.Count = 1
	erns := pinger.Run()
	if erns != nil {
		return "", erns
	}
	stats := pinger.Statistics().AvgRtt.String()
	return stats, nil
}

func IcmpPingConcurrent(hostList []string, pingTime int) (map[string][][]string, map[string][]error) {
	var comms = make(chan []string)
	var rawDataFinal = make(map[string][][]string)
	var errors = make(map[string][]error)
	for _, host := range hostList {
		go func(hostF string) {
			pinger, err := probing.NewPinger(hostF)
			pinger.SetPrivileged(true)
			if err != nil {
				errors[hostF] = append(errors[hostF], err)
			}
			go func() {
				erns := pinger.Run()
				if erns != nil {
					errors[hostF] = append(errors[hostF], erns)
				}
			}()
			go func() {
				for {
					stt := pinger.Statistics()
					stats := []string{
						hostF,
						stt.AvgRtt.String(),
						stt.MaxRtt.String(),
						stt.MinRtt.String(),
						fmt.Sprintf("%v", stt.PacketsSent),
						fmt.Sprintf("%v", stt.PacketsRecv),
						fmt.Sprintf("%v", stt.PacketLoss),
					}
					comms <- stats
					time.Sleep(time.Second / 20)
				}
			}()
		}(host)
	}
	go func() {
		for stats := range comms {
			rawDataFinal[stats[0]] = [][]string{stats}
		}
	}()
	time.Sleep(time.Second * time.Duration(pingTime))
	return rawDataFinal, errors
}

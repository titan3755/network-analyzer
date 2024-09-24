package utils

import (
	probing "github.com/prometheus-community/pro-bing"
	"strconv"
)

// ICMP_Ping is a function that sends an ICMP echo request to the specified host

func ICMP_Ping(host string) (latency string, er error) {
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

func ICMP_Ping_Concurrent(host string, commsChannel chan []string) {
	var running bool = true
	pinger, err := probing.NewPinger(host)
	pinger.SetPrivileged(true)
	if err != nil {
		return
	}
	go func() {
		erns := pinger.Run()
		if erns != nil {
			commsChannel <- []string{"stop"}
			return
		}
	}()
	for running {
		stats := pinger.Statistics()
		strStatList := []string{
			stats.AvgRtt.String(),
			stats.MaxRtt.String(),
			stats.MinRtt.String(),
			strconv.FormatInt(int64(stats.PacketsSent), 10),
			strconv.FormatInt(int64(stats.PacketsRecv), 10),
			strconv.FormatInt(int64(stats.PacketLoss), 10),
			stats.IPAddr.String(),
		}
		commsChannel <- strStatList
		// msg := <- commsChannel
		// if msg[0] == "stop" {
		// 	running = false
		// }
	}
}


package utils

import (
	probing "github.com/prometheus-community/pro-bing"
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
	stats := pinger.Statistics().AvgRtt
	return stats.String(), nil
}



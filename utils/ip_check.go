package utils

import (
	"strconv"
	"strings"
)

// CheckIfValidIPv4 is a function that checks if the IP address is valid or not

func CheckIfValidIPv4(ip string) bool {
	//check for spaces
	if strings.Contains(ip, " ") {
		return false
	}
	//check for invalid characters
	if strings.ContainsAny(ip, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+-=[]{}|;',<>?/") {
		return false
	}
	//check for invalid length
	if len(ip) > 21 || len(ip) < 9 {
		return false
	}
	//check for invalid IP address
	if strings.Count(ip, ".") != 3 {
		return false
	}
	//check for invalid IP address
	if !strings.Contains(ip, ":") {
		return false
	}
	// check if port number is within logical range
	if strings.Contains(ip, ":") {
		port := strings.Split(ip, ":")[1]
		num_port, er := strconv.Atoi(port)
		if er != nil {
			return false
		}
		// check if port situated at the end of the IP address
		if strings.Contains(port, ".") {
			return false
		}
		// check if port is within logical range
		if num_port < 1 || num_port > 65535 {
			return false
		}
	}
	return true
}
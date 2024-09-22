package utils

import (
	"strings"
)

// ConvListOfIPToFileFormat is a function that converts a list of IP addresses to a string format

func ConvListOfIPToFileFormat(ipList []string) string {
	var ipString string
	for _, ip := range ipList {
		ipString += ip + ";"
	}
	return ipString
}

// ConvFileFormatToListOfIP is a function that converts a string format of IP addresses to a list

func ConvFileFormatToListOfIP(ipString string) []string {
	return strings.Split(ipString, ";")
}
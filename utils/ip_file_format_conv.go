package utils

import (
	"strings"
)

func ConvListOfIPToFileFormat(ipList []string) string {
	var ipString string
	for _, ip := range ipList {
		ipString += ip + ";"
	}
	return ipString
}

func ConvFileFormatToListOfIP(ipString string) []string {
	return strings.Split(ipString, ";")
}
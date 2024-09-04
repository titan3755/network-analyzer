package utils

func ConvListOfIPToFileFormat(ipList []string) string {
	var ipString string
	for _, ip := range ipList {
		ipString += ip + ";"
	}
	return ipString
}

func ConvFileFormatToListOfIP(ipString string) []string {
	var ipList []string
	for _, ip := range ipString {
		if ip == ';' {
			ipList = append(ipList, string(ip))
		}
	}
	return ipList
}
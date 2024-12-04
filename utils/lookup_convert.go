package utils

func ConvertListOfHostsToIPs(hosts []string) (map[string][]string, []error) {
	var ips = make(map[string][]string)
	var errors []error
	for _, host := range hosts {
		ip, err := LookupHostIP(host)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		ips[host] = ip
	}
	return ips, errors
}

//func ConvertSingleHostToIPs(host string) ([]string, error) {
//	return LookupHostIP(host)
//}
//
//func ConvertListOfIpsToHosts(ips []string) (map[string][]string, []error) {
//	var hosts = make(map[string][]string)
//	var errors []error
//	for _, ip := range ips {
//		host, err := LookupAddrHost(ip)
//		if err != nil {
//			errors = append(errors, err)
//			continue
//		}
//		hosts[ip] = host
//	}
//	return hosts, errors
//}
//
//func ConvertSingleIPToHosts(ip string) ([]string, error) {
//	return LookupAddrHost(ip)
//}

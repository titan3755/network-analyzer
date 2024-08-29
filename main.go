package main

import "fmt"

// Server IP list
var s_ip []string = make([]string, 0)

// Functions -->

func addServerIP(ip string)  {
	s_ip = append(s_ip, ip)
}

func printServerIP()  {
	for _, ip := range s_ip {
		fmt.Println(ip)
	}
}

// Main function

func main()  {
	fmt.Println("Hello, World!")
}
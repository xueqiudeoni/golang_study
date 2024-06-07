package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	argument := os.Args
	//ip := net.ParseIP(argument[1])
	//if ip == nil {
	//	adds, err := net.LookupHost(argument[1])
	//	if err == nil {
	//		for _, i := range adds {
	//			fmt.Println("sigleip", i)
	//		}
	//	}
	//} else {
	//	hosts, _ := net.LookupAddr(argument[1])
	//	for _, host := range hosts {
	//		fmt.Println("host:", host)
	//	}
	//}
	//NSs, _ := net.LookupNS(argument[1])
	//for _, s := range NSs {
	//	fmt.Println(s, s.Host)
	//}
	MXs, _ := net.LookupMX(argument[1])
	for _, s := range MXs {
		fmt.Println(s, s.Host, s.Pref)
	}

}

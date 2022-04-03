package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func Port(Hostname, proto string, numberPort int) bool {
	PortAddr := Hostname + ":" + strconv.Itoa(numberPort)
	conn, err := net.DialTimeout(proto, PortAddr, 60*time.Second)
	if CheckErr(err) {
		fmt.Println(PortAddr + "pasive")
		return false
	}
	defer conn.Close()

	return true
}
func CheckErr(err error) bool {
	return err != nil
}

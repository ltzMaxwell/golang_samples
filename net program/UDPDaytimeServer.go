/* UDPDaytimeServer
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	/**
	 * address to connect is xxxx:1200
	 * @type {[type]}
	 */
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	fmt.Println("udpAddr is ", udpAddr)
	checkError(err)
	/**
	 * connection of UDP do not need 'accept'
	 * @type {[type]}
	 */
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	fmt.Println("handleClient")
	var buf [512]byte
	/**
	 * block here
	 * @type {[type]}
	 */
	_, addr, err := conn.ReadFromUDP(buf[0:])

	fmt.Println("addr read from udp is :", addr)

	if err != nil {
		return
	}

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
	fmt.Println("WriteToUDP")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

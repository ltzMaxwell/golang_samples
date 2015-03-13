/* DaytimeServer TCP
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"
	/**
	 * resolve address
	 * @type {[type]}
	 */
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	/**
	 * get listner
	 * @type {[type]}
	 */
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		/**
		 * get connection from listener
		 * @type {[type]}
		 */
		conn, err := listener.Accept()
		fmt.Println("connectd to server")
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
		fmt.Println("conn closed")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

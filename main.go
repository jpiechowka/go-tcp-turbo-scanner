package main

import (
	"fmt"
	"github.com/jpiechowka/go-tcp-turbo-scanner/scanner"
)

func main() {
	// TODO, min port, max port and host from command line

	host := "localhost"

	doneChan := make(chan struct{})
	defer close(doneChan)

	for tcpPortState := range scanner.ScanTCPPortsRange(doneChan, host, 0, 65535) {
		if tcpPortState.IsOpen {
			openPortMsg := fmt.Sprintf("[+] %s:%d port open", host, tcpPortState.PortNumber)
			fmt.Println(openPortMsg)
		}
	}
}

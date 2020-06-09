package main

import (
	"fmt"
	"github.com/jpiechowka/go-tcp-turbo-scanner/scanner"
)

func main() {
	// TODO, min port, max port and host from command line

	doneChan := make(chan struct{})
	defer close(doneChan)

	for tcpPortState := range scanner.ScanTCPPortsRange(doneChan, 0, 65535) {
		fmt.Println(tcpPortState.PortNumber)
	}
}

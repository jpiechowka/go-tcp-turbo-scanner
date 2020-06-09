package scanner

import (
	"sync"
)

const portNumberGeneratorBufferSize = 50
const portScanResultsBufferSize = 50

type TCPPortState struct {
	PortNumber int
	IsOpen     bool
}

func ScanTCPPortsRange(done <-chan struct{}, minPort int, maxPort int) <-chan TCPPortState {
	portsToScanChan := portNumberGenerator(done, minPort, maxPort)
	return scanTCPPortsRange(done, portsToScanChan)
}

func portNumberGenerator(done <-chan struct{}, minPort int, maxPort int) <-chan int {
	portNumberChan := make(chan int, portNumberGeneratorBufferSize)

	go func() {
		defer close(portNumberChan)

		for i := minPort; i <= maxPort; i++ {
			select {
			case <-done:
				return
			case portNumberChan <- i:
			}
		}
	}()

	return portNumberChan
}

func scanTCPPortsRange(done <-chan struct{}, portsChan <-chan int) <-chan TCPPortState {
	tcpPortScanResultChan := make(chan TCPPortState, portScanResultsBufferSize)

	go func() {
		defer close(tcpPortScanResultChan)

		wg := sync.WaitGroup{}

		for port := range portsChan {
			select {
			case <-done:
				return
			default:
				wg.Add(1)

				go func(tcpPort int) {
					defer wg.Done()
					tcpPortScanResultChan <- scanSingleTCPPort(tcpPort)
				}(port)
			}
		}

		wg.Wait()
	}()

	return tcpPortScanResultChan
}

func scanSingleTCPPort(port int) TCPPortState {
	// TODO Implement scanner

	tcpPortState := TCPPortState{
		PortNumber: port,
		IsOpen:     false,
	}

	return tcpPortState
}

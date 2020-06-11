package scanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

const portNumberGeneratorBufferSize = 100
const portScanResultsBufferSize = 100

type TCPPortState struct {
	PortNumber int
	IsOpen     bool
}

func ScanTCPPortsRange(done <-chan struct{}, host string, minPort int, maxPort int, maxConcurrency int, shouldPrintVerbose bool) <-chan TCPPortState {
	portsToScanChan := portNumberGenerator(done, minPort, maxPort)
	return scanTCPPortsRange(done, host, portsToScanChan, maxConcurrency, shouldPrintVerbose)
}

func scanTCPPortsRange(done <-chan struct{}, host string, portsChan <-chan int, maxConcurrency int, shouldPrintVerbose bool) <-chan TCPPortState {
	tcpPortScanResultChan := make(chan TCPPortState, portScanResultsBufferSize)
	concurrencyGuard := make(chan struct{}, maxConcurrency)

	go func() {
		defer close(tcpPortScanResultChan)
		defer close(concurrencyGuard)

		wg := sync.WaitGroup{}

		for port := range portsChan {
			select {
			case <-done:
				return
			default:
				wg.Add(1)
				concurrencyGuard <- struct{}{} // Will block if channel is filled

				go func(tcpPort int) {
					defer wg.Done()
					defer func() {
						<-concurrencyGuard
					}()
					tcpPortScanResultChan <- scanSingleTCPPort(host, tcpPort, shouldPrintVerbose)
				}(port)
			}
		}

		wg.Wait()
	}()

	return tcpPortScanResultChan
}

func scanSingleTCPPort(host string, port int, shouldPrintVerbose bool) TCPPortState {
	address := net.JoinHostPort(host, strconv.Itoa(port))

	if shouldPrintVerbose {
		fmt.Println("[V] Scanning: " + address)
	}

	tcpConnection, tcpDialError := net.Dial("tcp", address)

	if tcpDialError == nil {
		tcpConnection.Close()
	}

	tcpPortState := TCPPortState{
		PortNumber: port,
		IsOpen:     tcpDialError == nil,
	}

	return tcpPortState
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

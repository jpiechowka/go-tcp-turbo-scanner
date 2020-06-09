package main

import (
	"flag"
	"fmt"
	"github.com/jpiechowka/go-tcp-turbo-scanner/scanner"
	"os"
	"time"
)

func main() {
	doneChan := make(chan struct{})
	defer close(doneChan)

	// Parse command line arguments
	// TODO Error checking
	// In case of zero default value, default value message for flag is not printed. It needs to be explicitly added.
	minPort := flag.Int("min-port", 0, "Minimum port number to scan, inclusive (defaults to 0)")
	maxPort := flag.Int("max-port", 65535, "Maximum port number to scan, inclusive")
	// TODO Implement changing number of used threads / goroutines
	threads := flag.Int("threads", 100, "Number of threads to use for scanning")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s: [optional flags] <hostname>\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Multiple hosts can scanned. Specify hosts separated by space after cli flags, for example:\n")
		fmt.Printf("%s --min-port 1111 --max-port 13337 host1 host2 host3\n\n", os.Args[0])
		fmt.Printf("Optional flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check argument errors, at least one host is required
	if len(os.Args) < 2 {
		fmt.Println("One or more hosts are required to run")
		fmt.Println("Specify hosts separated by space after cli flags, for example:")
		fmt.Printf("%s --min-port 1111 --max-port 13337 host1 host2 host3\n", os.Args[0])
		os.Exit(1)
	}

	hosts := flag.Args()

	startTime := time.Now()

	for idx, host := range hosts {
		// TODO host error checking

		allOpenPorts := ""

		fmt.Printf("Starting [%d/%d] scan of %s at %s\n", idx+1, len(hosts), host, time.Now().String())
		fmt.Printf("Starting port: %d, Max port: %d, threads: %d\n\n", *minPort, *maxPort, *threads)

		for tcpPortState := range scanner.ScanTCPPortsRange(doneChan, host, *minPort, *maxPort) {
			if tcpPortState.IsOpen {
				allOpenPorts = allOpenPorts + fmt.Sprintf(" %d", tcpPortState.PortNumber)
				fmt.Printf("[+] %s:%d port open\n", host, tcpPortState.PortNumber)
			}
		}

		fmt.Printf("[+] %s all open ports:%s\n\n", host, allOpenPorts)
	}

	fmt.Printf("Finished scanning in %s", time.Since(startTime))
}

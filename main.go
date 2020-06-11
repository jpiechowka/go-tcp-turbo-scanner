package main

import (
	"flag"
	"fmt"
	"github.com/jpiechowka/go-tcp-turbo-scanner/scanner"
	"net"
	"os"
	"time"
)

func main() {
	doneChan := make(chan struct{})
	defer close(doneChan)

	// Parse command line arguments
	// In case of zero default value, default value message for flag is not printed. It needs to be explicitly added.
	minPort := flag.Int("min-port", 0, "Minimum port number to scan, inclusive (default 0)")
	maxPort := flag.Int("max-port", 65535, "Maximum port number to scan, inclusive")
	threads := flag.Int("threads", 20, "Number of threads to use for scanning")
	verbose := flag.Bool("verbose", false, "Enable verbose output")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\nUsage of %s: [optional flags] <hostname>\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Multiple hosts can scanned. Specify hosts separated by space after cli flags, for example:\n")
		fmt.Printf("%s --min-port 1111 --max-port 13337 host1 host2 host3\n\n", os.Args[0])
		fmt.Printf("Optional flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check min-port or/and max-port values if necessary
	if *minPort < 0 || *minPort >= 65534 {
		fmt.Println("--min-port value cannot be smaller than 0 and larger than 65534")
		os.Exit(1)
	} else if *minPort >= *maxPort {
		fmt.Println("--min-port value cannot be larger or equal than --max-port value")
		os.Exit(1)
	}

	if *maxPort < 1 || *maxPort > 65535 {
		fmt.Println("--max-port value cannot be smaller than 1 and larger than 65535")
		os.Exit(1)
	} else if *maxPort <= *minPort {
		fmt.Println("--max-port value cannot be smaller or equal than --min-port value")
		os.Exit(1)
	}

	//Check threads flag for correctness
	if *threads < 1 {
		fmt.Println("--threads value cannot be smaller than 1")
		os.Exit(1)
	}

	hosts := flag.Args()

	// Check argument errors, at least one host is required
	if len(hosts) < 1 {
		fmt.Println("One or more hosts are required to run")
		fmt.Println("Specify hosts separated by space after cli flags, for example:")
		fmt.Printf("%s --min-port 1111 --max-port 13337 host1 host2 host3\n", os.Args[0])
		os.Exit(1)
	}

	// Check validity of hosts
	if !areHostsValid(hosts) {
		fmt.Println("Provided host / hosts are invalid")
		os.Exit(1)
	}

	startTime := time.Now()

	for idx, host := range hosts {
		allOpenPorts := ""

		fmt.Printf("\nStarting [%d/%d] scan of %s at %s\n", idx+1, len(hosts), host, time.Now().String())
		fmt.Printf("Starting port: %d, Max port: %d, threads: %d\n\n", *minPort, *maxPort, *threads)

		for tcpPortState := range scanner.ScanTCPPortsRange(doneChan, host, *minPort, *maxPort, *threads, *verbose) {
			if tcpPortState.IsOpen {
				allOpenPorts = allOpenPorts + fmt.Sprintf(" %d", tcpPortState.PortNumber)
				fmt.Printf("[+] %s:%d port open\n", host, tcpPortState.PortNumber)
			}
		}

		if len(allOpenPorts) == 0 {
			fmt.Printf("\n[+] %s there are no open ports\n", host)
		} else {
			fmt.Printf("\n[+] %s all open ports:%s\n", host, allOpenPorts)
		}
	}

	fmt.Printf("\nFinished scanning %d hosts in %s", len(hosts), time.Since(startTime))
}

func areHostsValid(hosts []string) bool {
	for _, host := range hosts {
		addr := net.ParseIP(host)

		if addr == nil {
			_, lookupErr := net.LookupHost(host)
			if lookupErr != nil {
				return false
			}
		}
	}

	return true
}

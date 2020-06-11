# Go TCP turbo scanner
A very small, fast, concurrent TCP scanner written in Go

### Usage
```
Usage of go-tcp-turbo-scanner: [optional flags] <hostname>
Multiple hosts can scanned. Specify hosts separated by space after cli flags, for example:
go-tcp-turbo-scanner --min-port 1111 --max-port 13337 host1 host2 host3

Optional flags:
  -max-port int
        Maximum port number to scan, inclusive (default 65535)
  -min-port int
        Minimum port number to scan, inclusive (default 0)
  -threads int
        Number of threads to use for scanning (default 20)
  -verbose
        Enable verbose output
```

### Example output
Example scan of ```scanme.nmap.org```
```
./go-tcp-turbo-scanner --max-port 1500 --threads 200 scanme.nmap.org

Starting [1/1] scan of scanme.nmap.org at 2020-06-11 02:20:10.5493539 +0200 CEST m=+0.005925101
Starting port: 0, Max port: 1500, threads: 200

[+] scanme.nmap.org:80 port open
[+] scanme.nmap.org:22 port open

[+] scanme.nmap.org all open ports: 80 22

Finished scanning 1 hosts in 26.9171835s
```

### Releases
Release binaries for every operating system can be found on the releases page:
https://github.com/jpiechowka/go-tcp-turbo-scanner/releases

### Automated builds status
![Build, test and create release from version tag](https://github.com/jpiechowka/go-tcp-turbo-scanner/workflows/Build,%20test%20and%20create%20release%20from%20version%20tag/badge.svg)

![Build and test from latest commit on master](https://github.com/jpiechowka/go-tcp-turbo-scanner/workflows/Build%20and%20test%20from%20latest%20commit%20on%20master/badge.svg)

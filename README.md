# Go TCP turbo scanner
A very small, fast, concurrent TCP scanner written in Go

### Usage
```
Usage of go-tcp-turbo-scanner-v0.1.0-win.exe: [optional flags] <hostname>
Multiple hosts can scanned. Specify hosts separated by space after cli flags, for example:
go-tcp-turbo-scanner-v0.1.0-win.exe --min-port 1111 --max-port 13337 host1 host2 host3

Optional flags:
  -max-port int
        Maximum port number to scan, inclusive (default 65535)
  -min-port int
        Minimum port number to scan, inclusive (default 0)
  -threads int
        Number of threads to use for scanning (default 20)
```

### Releases
Release binaries for every operating system can be found on the releases page:
https://github.com/jpiechowka/go-tcp-turbo-scanner/releases

### Automated builds status
![Build, test and create release from version tag](https://github.com/jpiechowka/go-tcp-turbo-scanner/workflows/Build,%20test%20and%20create%20release%20from%20version%20tag/badge.svg)

![Build and test from latest commit on master](https://github.com/jpiechowka/go-tcp-turbo-scanner/workflows/Build%20and%20test%20from%20latest%20commit%20on%20master/badge.svg)

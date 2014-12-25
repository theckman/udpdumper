# UDPDumper
[![TravisCI Build Status](https://img.shields.io/travis/theckman/udpdumper/master.svg?style=flat)](https://travis-ci.org/theckman/udpdumper)
[![GoDoc](https://img.shields.io/badge/udpdumper-GoDoc-blue.svg?style=flat)](https://godoc.org/github.com/theckman/udpdumper/dumper)

UDPDumper is a small tool for printing all UDP communications to a single port.
This is meant for testing thing like statsd emissions, or anything else that
uses UDP.

## Installation
From the `master` branch:
```
go install github.com/theckman/udpdumper
```

## Usage
```
# host defaults to 127.0.0.1 and port defaults to 8125
udpdumper --host 127.0.0.2 --port 8130
random UDP traffic<EOF>
more random UDP traffic<EOF>

```

## License
UDPDumper is released under the BSD 3-Clause License. See the `LICENSE` file for
the full contents of the license.

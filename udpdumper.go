// Copyright 2014 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

type args struct {
	Host string `short:"h" long:"host" default:"127.0.0.1" description:"IP to bind to"`
	Port uint16 `short:"p" long:"port" default:"8125" description:"UDP port to bind to"`
}

func (a *args) parse() {
	parser := flags.NewParser(a, flags.HelpFlag|flags.PassDoubleDash)

	_, err := parser.Parse()

	// if parsing bombed and it's not the help message
	// we need to print the error
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", err.Error())
			os.Exit(0)
		}
	}
}

func NewUDPListener(host string, port uint16) *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%d", host, port))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return l
}

func Printer(listener *net.UDPConn, out chan<- []byte) {
	defer func() {
		_ = recover()
	}()

	defer close(out)

	for {
		b := make([]byte, 8192)

		_, err := listener.Read(b)

		if err != nil {
			close(out)
			break
		}

		out <- bytes.Trim(b, "\x00")
	}
}

func main() {
	var (
		a  = &args{}
		ch = make(chan []byte, 8)
	)

	a.parse()

	l := NewUDPListener(a.Host, a.Port)

	go Printer(l, ch)

	for v := range ch {
		fmt.Printf("%v<EOF>\n", string(v))
	}
}

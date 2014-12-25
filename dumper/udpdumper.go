package udpdumper

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

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

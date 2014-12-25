package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/theckman/udpdumper/dumper"
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

func main() {
	var (
		a  = &args{}
		ch = make(chan []byte, 8)
	)

	a.parse()

	l := udpdumper.NewUDPListener(a.Host, a.Port)

	go udpdumper.Printer(l, ch)

	for v := range ch {
		fmt.Printf("%v<EOF>\n", string(v))
	}
}

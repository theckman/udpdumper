package udpdumper_test

import (
	"net"
	"testing"

	"github.com/theckman/udpdumper/dumper"
	. "gopkg.in/check.v1"
)

const (
	HOST        = "127.0.0.1"
	PORT uint16 = 8125
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct {
	out      chan []byte
	listener *net.UDPConn
	writer   *net.UDPConn
}

var _ = Suite(&TestSuite{})

func (s *TestSuite) SetUpSuite(c *C) {
	s.out = make(chan []byte)
	s.listener = udpdumper.NewUDPListener(HOST, PORT)

	w, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP(HOST), Port: int(PORT)})

	c.Assert(err, IsNil)

	s.writer = w
}

func (s *TestSuite) TearDownSuite(c *C) {
	s.writer.Close()
	s.listener.Close()
}

func (s *TestSuite) TestNewUDPListener(c *C) {
	l := udpdumper.NewUDPListener("127.0.0.1", 8130)

	c.Assert(l.LocalAddr().String(), Equals, "127.0.0.1:8130")
}

func (s *TestSuite) TestPrinter(c *C) {
	go udpdumper.Printer(s.listener, s.out)

	s.writer.Write([]byte("hello there! -- 1234\n"))

	// it should return what has been written
	r, ok := <-s.out

	c.Assert(ok, Equals, true)
	c.Assert(string(r), Equals, "hello there! -- 1234\n")

	// when the listener is closed
	// it should close the return channel
	s.listener.Close()

	r, ok = <-s.out

	c.Assert(ok, Equals, false)
	c.Assert(r, IsNil)
}

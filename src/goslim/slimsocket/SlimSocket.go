package slimSocket

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

var ip_and_port string

func init() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	ip_and_port = ":" + os.Args[1]
}

type slimSocket struct {
	conn net.Conn
}

var oneSingleton *slimSocket

func Instance() *slimSocket {
	if oneSingleton == nil {
		oneSingleton = slimSocket_Create()
	}
	return oneSingleton
}

func slimSocket_Create() *slimSocket {
	return new(slimSocket)
}

func (self *slimSocket) Listen() {
	listener, err := net.Listen("tcp", ip_and_port)
	if err != nil {
		return
	}

	conn, err := listener.Accept()
	self.checkError(err)

	self.conn = conn

	connFrom := self.conn.RemoteAddr().String()
	log.Println("Connection from: ", connFrom)
}

func (self *slimSocket) SendMsg(msg string) {
	_, err := self.conn.Write([]byte(msg))
	self.checkError(err)
}

func (self *slimSocket) checkError(err error) {
	if err != nil {
		panic("ERROR: " + err.Error()) // terminate program
	}
}

func (self *slimSocket) receive(readLen int) ([]byte, error) {
	var buf []byte = make([]byte, readLen)

	receivedLen := 0
	for {
		if curLen, err := self.conn.Read(buf[receivedLen:readLen]); err != nil {
			return buf, err
		} else {
			receivedLen += curLen
		}

		if receivedLen >= readLen {
			break
		}
	}

	return buf, nil
}

func (self *slimSocket) readSize() int {
	size, err := self.receive(6)
	if err != nil {
		return -1
	}
	colon, err := self.receive(1)
	if (err != nil) || (colon[0] != ':') {
		return -1
	}

	length, err := strconv.Atoi(string(size))
	if err != nil {
		return -1
	}

	return length
}

func (self *slimSocket) ReceiveMsg() (string, error) {
	size_i := self.readSize()
	if size_i <= 0 {
		return "", errors.New("read size exception")
	}

	numbytes, err := self.receive(size_i)
	if err != nil || len(numbytes) != size_i {
		log.Println("did not receive right number of bytes.  %d expected but received %d\n", size_i, numbytes)
		return "", errors.New("receive exception")
	}

	return string(numbytes), nil
}

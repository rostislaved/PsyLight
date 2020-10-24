package port

import (
	"github.com/jacobsa/go-serial/serial"
	"io"
	"log"
)

//defer profile.Start().Stop()

type port struct {
	Conn io.ReadWriteCloser
}

func New() *port {
	options := serial.OpenOptions{
		PortName: "/dev/ttyUSB0",
		//BaudRate:        115200,
		BaudRate:        1000000,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}

	// Open the port.
	conn, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	return &port{Conn: conn}
}

func (p *port) Write(b []byte) {
	_, err := p.Conn.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}
}

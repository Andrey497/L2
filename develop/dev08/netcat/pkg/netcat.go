package pkg

import (
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"time"
)

type netcat struct {
	port     string
	host     string
	protocol string
}

func NewNetcat() *netcat {
	return &netcat{port: "", host: "", protocol: ""}
}

func connect(port, host, protocol string) (net.Conn, error) {
	conn, err := net.Dial(protocol, host+":"+port)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
func (n *netcat) Start() error {
	err := n.switchFlags()
	if err != nil {
		return err
	}

	if n.protocol == "tcp" {
		go ServerTcp(n.port)
		time.Sleep(2 * time.Second)
		clientTcp(n.port, n.host, n.protocol)
	} else if n.protocol == "udp" {
		go ServerUdp(n.port)
		time.Sleep(2 * time.Second)
		clientUdp(n.port, n.host, n.protocol)
	}

	return nil

}

func (n *netcat) switchFlags() error {
	port := flag.String("port", "8080", "протокол соединения ")
	host := flag.String("host", "localhost", "хост соединения ")
	protocol := flag.String("protocol", "tcp", "протокол соединения")
	flag.Parse()
	if *protocol != "tcp" && *protocol != "udp" {
		return errors.New("bad protocol")
	}
	n.port = *port
	n.host = *host
	n.protocol = *protocol
	return nil
}

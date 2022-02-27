package pkg

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
	conn, err := connect(n.port, n.host, n.protocol)
	if err != nil {
		return err
	}

	fmt.Println("Start server")
	go copyTo(os.Stdout, conn) // читаем из сокета в stdout
	copyTo(conn, os.Stdin)     // пишем в сокет из stdin
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

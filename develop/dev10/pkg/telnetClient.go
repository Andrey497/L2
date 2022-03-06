package pkg

import (
	"fmt"
	"io"
	"net"
	"time"
)

type telnetClient struct {
	host    string
	port    string
	timeOut time.Duration
	conn    net.Conn
	reader  io.Reader
	writer  io.Writer
}

func NewClient(host, port string, timeOut time.Duration, reader io.Reader, writer io.Writer) *telnetClient {
	return &telnetClient{
		host: host, port: port, timeOut: timeOut, reader: reader, writer: writer,
	}
}

func (t *telnetClient) BuildConnect() error {
	var err error
	t.conn, err = net.DialTimeout("tcp", t.host+":"+t.port, t.timeOut)
	return err
}

func (t *telnetClient) Send() error {
	fmt.Scan()
	data := []byte{}
	_, err := t.reader.Read(data)
	if err != nil {
		return err
	}
	if _, err := t.conn.Write(data); err != nil {
		return err
	}
	return nil
}

func (t *telnetClient) Receive() error {
	buff := make([]byte, 1024)
	n, err := t.conn.Read(buff)
	_, err = t.writer.Write(buff[0:n])
	if err != nil {
		return err
	}

	return nil
}
func (t *telnetClient) Close() error {
	return t.conn.Close()
}

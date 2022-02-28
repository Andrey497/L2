package pkg

import (
	"fmt"
	"net"
)

func ServerUdp(port string) {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Println("resolveaddr err", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listen udp err", err)
		return
	}
	var n int
	go func() {
		for {
			buf := make([]byte, 4096)
			n, addr, err = conn.ReadFromUDP(buf)
			fmt.Println("return: ", string(buf[:n]))
		}
	}()
}

// обработка подключения

func clientUdp(port, host, protocol string) error {
	fmt.Println("client add")
	conn, error := connectUdp(host+":"+port, protocol)
	if error != nil {
		fmt.Println(error)
		return error
	}
	go readUdp(conn)
	writeUdp(conn)
	defer conn.Close()
	return nil
}

func readUdp(conn net.Conn) error {
	for {
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			return err
		}
		fmt.Println(string(buff[0:n]))
	}
	return nil
}
func writeUdp(conn net.Conn) error {
	var source string
	for {
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
		}
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func connectUdp(host, protocol string) (net.Conn, error) {
	conn, err := net.Dial(protocol, host)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

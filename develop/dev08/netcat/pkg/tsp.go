package pkg

import (
	"fmt"
	"net"
)

func ServerTcp(port string) {
	listener, err := net.Listen("tcp", "127.0.0.1"+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn) // запускаем горутину для обработки запроса
	}

}

// обработка подключения
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// считываем полученные в запросе данные
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		conn.Write([]byte("return:" + string(input)))
	}
}
func clientTcp(port, host, protocol string) error {
	fmt.Println("client add")
	conn, error := connectTCP(port, host, protocol)
	if error != nil {
		return error
	}
	go readTcp(conn)
	writeTcp(conn)
	defer conn.Close()
	return nil
}

func readTcp(conn net.Conn) error {
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
func writeTcp(conn net.Conn) error {
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

func connectTCP(port, host, protocol string) (net.Conn, error) {
	conn, err := net.Dial(protocol, host+":"+port)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

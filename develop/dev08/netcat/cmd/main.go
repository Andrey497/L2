package main

import (
	"fmt"
	"net"
	"netcat/pkg"
)

/*Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func server() {
	message := "Hello, I am a server" // отправляемое сообщение
	listener, err := net.Listen("tcp", ":8080")

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
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
func main() {
	go server()
	netcat := pkg.NewNetcat()
	err := netcat.Start()
	if err != nil {
		panic(err)
	}
}

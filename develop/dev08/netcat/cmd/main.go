package main

import (
	"fmt"
	"netcat/pkg"
)

/*Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(string([]byte{117, 110, 100, 101, 102, 105, 110, 101, 100}))
	netcat := pkg.NewNetcat()
	err := netcat.Start()
	if err != nil {
		panic(err)
	}
}

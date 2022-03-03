package main

import (
	"factory/pkg"
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

func main() {
	tr1 := pkg.GetTransportByType("boat")
	tr2 := pkg.GetTransportByType("plane")
	tr3 := pkg.GetTransportByType("car")
	fmt.Println(tr1.GetType())
	fmt.Println(tr2.GetType())
	fmt.Println(tr3.GetType())
}

package main

import (
	"fmt"
	"pattern_fasad/pkg"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
func main() {
	w := pkg.CreateWallet(1, "testAccount", "+7988324412", "asda@1", 1)
	err := w.AddMoney("+7988324412", "asda@1", 100)
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"log"
	"state/pkg"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
func main() {
	atm := pkg.NewAtm(100)
	err := atm.PowerOn()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = atm.GiveMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = atm.AddMoney(100)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = atm.PowerOff()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

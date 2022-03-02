package main

import (
	"builder/pkg"
	"fmt"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
func main() {
	trucBuilder := pkg.GetCarBuilder("truc")
	director := pkg.NewCarDirector(trucBuilder)
	truc := director.BuildCar()
	fmt.Println(*truc)
}

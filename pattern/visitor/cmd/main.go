package main

import "visitor/pkg"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/
func main() {
	car := pkg.NewCar()
	ship := pkg.NewShip()
	plane := pkg.NewPlane()
	move := pkg.NewMove()
	car.Accept(&move)
	ship.Accept(&move)
	plane.Accept(&move)
}

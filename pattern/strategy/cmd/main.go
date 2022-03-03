package main

import "strategy/pkg"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

func main() {
	user := pkg.NewUser("testUser")

	bus := pkg.NewBus()
	walk := pkg.NewWalking()
	car := pkg.NewCar()
	taxy := pkg.NewTax()

	user.SetRout("Aviamotornaya 1", "Arbat 2")

	user.SetTypeMoving(bus)
	user.Moving()
	user.SetTypeMoving(walk)
	user.Moving()
	user.SetTypeMoving(car)
	user.Moving()
	user.SetTypeMoving(taxy)
	user.Moving()
}

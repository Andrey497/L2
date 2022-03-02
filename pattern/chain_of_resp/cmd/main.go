package main

import "chain_of_resp/pkg"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
func main() {
	account1 := pkg.Account{"login", "password", true}
	account2 := pkg.Account{"login", "password", false}
	checkAccount1 := pkg.CheckAccount{Login: "login", Password: "password"}
	checkAccount2 := pkg.CheckAccount{Login: "login", Password: "password"}
	checkIsAdmin1 := pkg.CheckIsAdministrator{IsAdministrator: true}
	checkIsAdmin2 := pkg.CheckIsAdministrator{IsAdministrator: false}
	removeTable1 := pkg.RemoveTable{}
	removeTable2 := pkg.RemoveTable{}
	checkAccount1.SetNext(&checkIsAdmin1)
	checkAccount2.SetNext(&checkIsAdmin2)
	checkIsAdmin1.SetNext(&removeTable1)
	checkIsAdmin2.SetNext(&removeTable2)
	checkAccount1.Execute(account1)
	checkAccount1.Execute(account2)
}

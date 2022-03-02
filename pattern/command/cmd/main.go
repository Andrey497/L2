package main

import "command/pkg"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
func main() {
	tv := &pkg.Tv{}
	commandOn := pkg.OnCommand{tv}
	commandOff := pkg.OffCommand{tv}
	commandCurrentChannel := pkg.CurrentChannelCommand{tv}
	commandNextChannel := pkg.NextChannelCommand{tv}
	commandPrevChannel := pkg.PrevChannelCommand{tv}
	commandOn.Execute()
	commandCurrentChannel.Execute()
	commandNextChannel.Execute()
	commandNextChannel.Execute()
	commandPrevChannel.Execute()
	commandOff.Execute()
}

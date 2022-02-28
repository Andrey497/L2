package main

import (
	"bufio"
	"fmt"
	"os"
	"shell/pkg"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
*/

func main() {
	fmt.Println("Hello my shell,to exit enter exit")
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		args := strings.Fields(txt)
		command := args[0]
		if command == "exit" {
			break
		}
		parameter := strings.Join(args[1:], " ")
		shell := pkg.NewShell(command, parameter)
		shell.Execute()
	}

}

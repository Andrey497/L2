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
		text := sc.Text()
		arrayPlain := strings.Split(text, "|")
		for _, txt := range arrayPlain {
			args := strings.Fields(txt)
			if len(args) != 0 {
				command := args[0]
				if command == "exit" {
					return
				}

				parameter := strings.Join(args[1:], " ")

				shell := pkg.NewShell(command, parameter)
				err := shell.Execute()

				if err != nil {
					fmt.Println(err)
				}
			}

		}

	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	var cmd *exec.Cmd
	cmd = exec.Command("ps")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		panic(err)
	}

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		panic(err)
	}

	if err = cmd.Wait(); err != nil {
		panic(err)
	}
	fmt.Print(string(data))

}

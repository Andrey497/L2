package main

import (
	"os"
	"strings"
	"wget/pkg"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	args := os.Args[1:] //получаем аргументы с терминала
	url := args[0]
	data := pkg.MakeRequest(url)
	nameFileArray := strings.Split(url, "/")
	nameFile := strings.Join(nameFileArray[len(nameFileArray)-3:], "_") + ".html"
	pkg.WriteDatInfile(nameFile, data)
}

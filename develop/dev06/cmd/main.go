package main

import (
	"bufio"
	"cut/cut"
	"io"
	"os"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func Read(input io.Reader) []string {
	result := []string{}
	in := bufio.NewScanner(input)
	for in.Scan() {
		string := in.Text()
		result = append(result, string)
	}
	return nil
}
func readFile(fileName string) []string {
	result := []string{}
	_, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		result = append(result, sc.Text())
	}
	defer file.Close()
	return result
}

func main() {
	args := os.Args[1:] //получаем аргументы с терминала
	if len(args) == 0 {
		panic("NO argument")
		return
	}
	fileName := args[len(args)-1] //имя файла
	lines := readFile(fileName)
	c, err := cut.InitCut(lines)
	if err != nil {
		panic(err)
	}
	c.Start()
	c.PrintResult()

}

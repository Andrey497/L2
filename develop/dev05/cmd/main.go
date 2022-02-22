/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"bufio"
	"greep/greep"
	"os"
)

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
	fileName := args[len(args)-1]     //имя файла
	targetString := args[len(args)-2] //имя файла
	lines := readFile(fileName)
	g, err := greep.InitGreep(lines, targetString)
	if err != nil {
		panic(err)
	}
	err = g.Start()
	if err != nil {
		panic(err)
	}
	g.GetResult()

}

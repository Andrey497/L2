package main

import (
	"bufio"
	"fmt"
	"os"
	"sorted/sorted"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
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
	fileName := args[0] //имя файла
	lines := readFile(fileName)
	s, err := sorted.InitSort(lines)
	if err != nil {
		panic(err)
	}
	s.SwitchFlags()

	err = s.Start()
	if err != nil {
		panic(err)
	}
	fmt.Println(fileName)
	result := s.GetResult()
	for _, val := range result {
		fmt.Println(val)
	}
	fmt.Println(s.GetIsSort())

}

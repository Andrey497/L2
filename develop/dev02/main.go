package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var (
	errInvalidString = errors.New("invalid string")
)

func UnpackageString(s string) (string, error) {
	runesBase := []rune(s)
	lengthString := len(runesBase)
	resultRunes := []rune{}
	i := 0
	var count int

	for i < lengthString {
		if string(runesBase[i]) == `\` { // if backSlash write the following character
			i++
			resultRunes = append(resultRunes, runesBase[i])
		} else if unicode.IsDigit(runesBase[i]) { //if is character is number
			count, _ = strconv.Atoi(string(runesBase[i]))
			if i == 0 ||
				(i < lengthString-1 && unicode.IsDigit(runesBase[i+1])) || //if following character is number
				(count < 1) {
				return "", errInvalidString
			} else { //if is character is number and not error
				endElem := resultRunes[len(resultRunes)-1]
				for j := 0; j < count-1; j++ {
					resultRunes = append(resultRunes, endElem)
				}
			}
		} else {
			resultRunes = append(resultRunes, runesBase[i])
		}
		i++

	}
	return string(resultRunes), nil
}
func main() {
	result, _ := UnpackageString("a4bc2d5e")
	if result == "aaaabccddddde" {
		fmt.Println(1)
		fmt.Println(result)
	}
	result, _ = UnpackageString("abcd")
	if result == "abcd" {
		fmt.Println(2)
		fmt.Println(result)
	}
	result, err := UnpackageString("45")
	if err != nil {
		fmt.Println(3)
		fmt.Println(err)
		fmt.Println(result)
	}
	result, err = UnpackageString(``)
	if result == "" {
		fmt.Println(4)
		fmt.Println(result)
	}
	result, err = UnpackageString(`qwe\4\5`)
	if result == "qwe45" {
		fmt.Println(5)
		fmt.Println(result)
	}
	result, err = UnpackageString(`qwe\45`)
	if result == `qwe44444` {
		fmt.Println(6)
		fmt.Println(result)
	}
	result, err = UnpackageString(`qwe\\5`)
	if result == `qwe\\\\\` {
		fmt.Println(7)
		fmt.Println(result)
	}
}

package main

import (
	"errors"
	"strconv"
	"unicode"
)

var (
	errInvalidString = errors.New("invalid string")
)

func UnPackageString(s string) (string, error) {
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

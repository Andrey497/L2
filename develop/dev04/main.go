package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

type sortRunes []rune

func createSortRunes(s string) sortRunes {
	var runes sortRunes
	runes = []rune(s)
	return runes
}

func (s *sortRunes) Len() int {
	return len(*s)
}

func (s *sortRunes) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *sortRunes) Less(i int, j int) bool {
	return (*s)[i] < (*s)[j]
}

func checkAnagram(s1 string, s2 string) bool {
	runes1 := createSortRunes(s1)
	runes2 := createSortRunes(s2)
	sort.Sort(&runes1)
	sort.Sort(&runes2)
	if string(runes1) == string(runes2) {
		return true
	}
	return false

}

func Anagrams(args *[]string) *map[string][]string {
	flagCheckAnagram := false
	result := make(map[string][]string)
	for _, item := range *args {
		lowerItem := strings.ToLower(item)
		for key, _ := range result {
			if checkAnagram(key, item) {
				flagCheckAnagram = true
				result[key] = append(result[key], lowerItem)
				break
			}
		}
		if !flagCheckAnagram {
			result[lowerItem] = []string{lowerItem}
		}
		flagCheckAnagram = false
	}

	for key, val := range result {
		if len(val) == 1 {
			delete(result, key)
			continue
		}
		sort.Strings(val)
		result[key] = val
	}

	return &result
}

func main() {
	slice := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(checkAnagram("тяпка", "пятак"))
	result := Anagrams(&slice)
	fmt.Println(*result)

}

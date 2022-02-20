package main

import (
	"anagram/Anagram"
	"fmt"
)

func main() {
	slice := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(Anagram.CheckAnagram("тяпка", "пятак"))
	result := Anagram.Anagrams(&slice)
	fmt.Println(*result)

}

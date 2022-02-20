package Anagram

import (
	"sort"
	"strings"
)

func CheckAnagram(s1 string, s2 string) bool {
	runes1 := CreateSortRunes(s1)
	runes2 := CreateSortRunes(s2)
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
			if CheckAnagram(key, item) {
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

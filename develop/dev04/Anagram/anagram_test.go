package Anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckAnagram(t *testing.T) {
	testTable := []struct {
		string1 string
		string2 string
		result  bool
	}{
		{
			string1: "тяпка",
			string2: "пятак",
			result:  true,
		},
		{
			string1: "",
			string2: "",
			result:  true,
		},
		{
			string1: "тряпка",
			string2: "торт",
			result:  false,
		},
		{
			string1: "стол",
			string2: "тост",
			result:  false,
		},
	}
	for _, testcase := range testTable {
		result := CheckAnagram(testcase.string1, testcase.string2)
		assert.Equal(t, testcase.result, result)
	}

}

func TestAnagrams(t *testing.T) {
	testTable := []struct {
		inputStrings []string
		expected     map[string][]string
	}{
		{
			inputStrings: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			expected:     map[string][]string{"листок": []string{"листок", "слиток", "столик"}, "пятак": []string{"пятак", "пятка", "тяпка"}},
		},
	}
	for _, testcase := range testTable {
		result := Anagrams(&testcase.inputStrings)
		assert.Equal(t, testcase.expected, *result)
	}

}

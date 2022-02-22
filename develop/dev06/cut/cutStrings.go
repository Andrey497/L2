package cut

import "strings"

func cutStrings(textLines []string, fields []int, delimiter string) []string {
	resultString := []string{}
	for _, val := range textLines {
		wordsInLines := []string{}
		if !strings.Contains(val, delimiter) {
			resultString = append(resultString, val)
			continue
		}
		words := strings.Split(val, delimiter)
		for _, item := range fields {
			if item < len(words)-1 {
				wordsInLines = append(wordsInLines, words[item])
			}
		}
		resultString = append(resultString, strings.Join(wordsInLines[:], delimiter))
	}
	return resultString
}

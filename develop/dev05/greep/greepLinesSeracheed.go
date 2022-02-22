package greep

import (
	"fmt"
	"sort"
)

func min(x1 int, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}

func greepLinesSearched(linesText []string, numberLine []int, after int, before int, numberLineFlag bool) []string {
	mapNumberString := make(map[int]struct{})
	keys := []int{}
	result := []string{}
	for _, val := range numberLine {
		lenBefore := min(val, before)
		lenAfter := min(after, len(linesText)-val-1)
		for i := val - lenBefore; i <= val+lenAfter; i++ {
			if _, ok := mapNumberString[i]; !ok {
				mapNumberString[i] = struct{}{}
				keys = append(keys, i)
			}

		}
	}
	sort.Ints(keys)
	for _, key := range keys {
		if numberLineFlag {
			result = append(result, fmt.Sprintf("%d: %s", key+1, linesText[key]))
		} else {
			result = append(result, linesText[key])

		}
	}
	return result
}

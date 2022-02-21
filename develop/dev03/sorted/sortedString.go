package sorted

import (
	"sort"
	"strings"
)

func sortedString(lines []string, reverse bool, spaceIgnore bool) []string {
	result := lines
	if spaceIgnore {
		for i, val := range result {
			result[i] = strings.TrimSpace(val)
		}
	}
	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(result)))
	} else {
		sort.Strings(result)
	}
	return result
}

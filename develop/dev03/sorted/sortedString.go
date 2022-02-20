package sorted

import "sort"

func sortedString(lines []string, reverse bool) []string {
	result := lines
	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(result)))
	} else {
		sort.Strings(result)
	}
	return result
}

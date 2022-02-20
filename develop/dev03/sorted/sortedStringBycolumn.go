package sorted

import (
	"sort"
)

func sortedStringByColumn(column int, lines []string, reverse bool) ([]string, error) {
	result := []string{}
	mapForSort := map[string][]string{}
	keys := []string{}
	for _, value := range lines {
		key, err := searchFieldFotSorted(value, column)
		if err != nil {
			return []string{}, err
		}
		mapForSort[key] = append(mapForSort[key], value)
		keys = append(keys, key)
	}
	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}
	for _, item := range keys {
		result = append(result, mapForSort[item]...)
	}
	return result, nil
}

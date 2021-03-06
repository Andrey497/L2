package sorted

import (
	"sort"
	"strings"
)

func sortedStringByColumn(column int, lines []string, reverse bool, spaseIgnor bool) ([]string, error) {
	result := []string{}
	mapForSort := map[string][]string{}
	keys := []string{}
	for _, value := range lines {
		key, err := searchFieldFotSorted(value, column)
		if err != nil {
			return []string{}, err
		}
		if _, ok := mapForSort[key]; !ok {
			if spaseIgnor {
				key = strings.TrimSpace(key)
			}
			keys = append(keys, key)
		}
		mapForSort[key] = append(mapForSort[key], value)
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

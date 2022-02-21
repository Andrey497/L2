package sorted

import (
	"sort"
	"strconv"
)

func sortedByNumber(column int, lines []string, reverse bool) ([]string, error) {
	result := []string{}
	mapForSort := map[int][]string{}
	keys := []int{}
	for _, value := range lines {
		key, err := searchFieldFotSorted(value, column)
		if err != nil {
			return []string{}, err
		}
		keyInt, err := strconv.Atoi(key)
		if err != nil {
			return []string{}, err
		}
		if _, ok := mapForSort[keyInt]; !ok {
			keys = append(keys, keyInt)
		}
		mapForSort[keyInt] = append(mapForSort[keyInt], value)

	}
	if reverse {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	} else {
		sort.Ints(keys)
	}
	for _, item := range keys {
		result = append(result, mapForSort[item]...)
	}
	return result, nil
}

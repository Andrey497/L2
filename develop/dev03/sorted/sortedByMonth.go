package sorted

import (
	"sort"
	"strings"
)

type Month []string

var month = map[string]int{
	"JAN": 1,
	"FAB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

func (m *Month) Len() int {
	return len(*m)
}
func (m *Month) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}
func (m *Month) Less(i int, j int) bool {
	return month[(*m)[i]] < month[(*m)[j]]
}

func sortedByMonth(column int, lines []string, reverse bool) ([]string, error) {
	result := []string{}
	notMonth := []string{}
	mapForSort := map[string][]string{}
	keys := Month{}
	for _, value := range lines {
		key, err := searchFieldFotSorted(value, column)
		if err != nil {
			return []string{}, err
		}

		key = strings.ToUpper(key)
		if _, ok := month[key]; !ok {
			notMonth = append(notMonth, value)
		}
		if _, ok := mapForSort[key]; !ok {
			keys = append(keys, key)
		}
		mapForSort[key] = append(mapForSort[key], value)

	}
	if reverse {
		sort.Sort(sort.Reverse(&keys))
	} else {
		sort.Sort(&keys)
	}
	for _, item := range keys {
		result = append(result, mapForSort[item]...)
	}
	result = append(result, notMonth...)
	return result, nil
}

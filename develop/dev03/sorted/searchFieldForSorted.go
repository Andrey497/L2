package sorted

import (
	"errors"
	"strings"
)

func searchFieldFotSorted(line string, column int) (string, error) {
	columns := strings.Fields(line)
	if len(columns) < column {
		return "", errors.New("index column out of range0")
	}
	return columns[column], nil
}

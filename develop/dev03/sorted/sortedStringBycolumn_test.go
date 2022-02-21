package sorted

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedColumn(t *testing.T) {
	testTable := []struct {
		lines       []string
		column      int
		reverse     bool
		spaceIgnore bool
		expected    []string
		err         error
	}{
		{
			lines:       []string{""},
			column:      2,
			reverse:     false,
			spaceIgnore: false,
			expected:    []string{},
			err:         errors.New("index column out of range0"),
		},
		{
			lines:       []string{"а а я", "а д с", "б б ф", "к д с"},
			column:      1,
			reverse:     false,
			spaceIgnore: false,
			expected:    []string{"а а я", "б б ф", "а д с", "к д с"},
			err:         nil,
		},
		{
			lines:       []string{"а а я", "а д с", "б б ф", "к д с"},
			column:      1,
			reverse:     true,
			spaceIgnore: true,
			expected:    []string{"а д с", "к д с", "б б ф", "а а я"},
			err:         nil,
		},
		{lines: []string{"а д с"},
			column:      1,
			reverse:     true,
			spaceIgnore: false,
			expected:    []string{"а д с"},
			err:         nil,
		},
	}

	for _, testcase := range testTable {
		result, err := sortedStringByColumn(testcase.column, testcase.lines, testcase.reverse, testcase.spaceIgnore)
		assert.Equal(t, testcase.expected, result)
		assert.Equal(t, testcase.err, err)
	}
}

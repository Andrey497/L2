package sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedByMonth(t *testing.T) {
	testTable := []struct {
		lines    []string
		column   int
		reverse  bool
		expected []string
		err      error
	}{
		{
			lines:    []string{"jun", "may", "aug"},
			column:   0,
			reverse:  false,
			expected: []string{"may", "jun", "aug"},
			err:      nil,
		},
		{
			lines:    []string{"jun", "may", "aug"},
			column:   0,
			reverse:  true,
			expected: []string{"aug", "jun", "may"},
			err:      nil,
		},
		{
			lines:    []string{"a jun", "a may", "a aug"},
			column:   1,
			reverse:  false,
			expected: []string{"a may", "a jun", "a aug"},
			err:      nil,
		},
		{
			lines:    []string{"a jun", "a may", "a aug"},
			column:   1,
			reverse:  true,
			expected: []string{"a aug", "a jun", "a may"},
			err:      nil,
		},
	}

	for _, testcase := range testTable {
		result, err := sortedByMonth(testcase.column, testcase.lines, testcase.reverse)
		assert.Equal(t, testcase.expected, result)
		assert.Equal(t, testcase.err, err)
	}
}

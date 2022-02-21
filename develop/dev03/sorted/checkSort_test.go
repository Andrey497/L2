package sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckSort(t *testing.T) {
	testTable := []struct {
		lines       []string
		sortedLines []string
		expected    bool
	}{
		{
			lines:       []string{"тест", "тест2"},
			sortedLines: []string{"тест", "тест2"},
			expected:    true,
		},
		{
			lines:       []string{"ямы", "ямы", "тест", "ямы"},
			sortedLines: []string{"ямы", "ямы", "тест", "ямы"},
			expected:    true,
		},
		{
			lines:       []string{"ямы"},
			sortedLines: []string{"ямы"},
			expected:    true,
		},
		{
			lines:       []string{},
			sortedLines: []string{},
			expected:    true,
		},
		{
			lines:       []string{"тест", "тест2"},
			sortedLines: []string{"тест2", "тест"},
			expected:    false,
		},
		{
			lines:       []string{"ямы f", "ямы f", "тест f", "ямы f"},
			sortedLines: []string{"ямы f", "ямы f", "ямы f", "тест f"},
			expected:    false,
		},
	}

	for _, testcase := range testTable {
		result := checkSort(testcase.lines, testcase.sortedLines)
		assert.Equal(t, testcase.expected, result)
	}
}

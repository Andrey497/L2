package greep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreepLinesSearched(t *testing.T) {
	testTable := []struct {
		linesString    []string
		numberLine     []int
		after          int
		before         int
		numberLineFlag bool
		expected       []string
	}{
		{
			linesString:    []string{"раз", "два", "три", "раз два три", "раз три четыре", "3 2 1"},
			numberLine:     []int{1, 3},
			after:          0,
			before:         0,
			numberLineFlag: false,
			expected:       []string{"два", "раз два три"},
		},
		{
			linesString:    []string{"раз", "два", "три", "раз два три", "раз три четыре", "3 2 1"},
			numberLine:     []int{1, 3},
			after:          1,
			before:         0,
			numberLineFlag: false,
			expected:       []string{"два", "три", "раз два три", "раз три четыре"},
		},
		{
			linesString:    []string{"раз", "два", "три", "раз два три", "раз три четыре", "3 2 1"},
			numberLine:     []int{1, 3},
			after:          1,
			before:         2,
			numberLineFlag: false,
			expected:       []string{"раз", "два", "три", "раз два три", "раз три четыре"},
		},
		{
			linesString:    []string{"раз", "два", "три", "раз два три", "раз три четыре", "3 2 1"},
			numberLine:     []int{1, 3},
			after:          1,
			before:         2,
			numberLineFlag: true,
			expected:       []string{"1: раз", "2: два", "3: три", "4: раз два три", "5: раз три четыре"},
		},
	}

	for _, testcase := range testTable {
		result := greepLinesSearched(testcase.linesString, testcase.numberLine, testcase.after, testcase.before, testcase.numberLineFlag)
		assert.Equal(t, testcase.expected, result)

	}
}

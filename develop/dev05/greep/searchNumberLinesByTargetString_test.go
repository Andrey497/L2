package greep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchFieldForSorted(t *testing.T) {
	testTable := []struct {
		linesString  []string
		targetString string
		ignoreCase   bool
		invert       bool
		fixed        bool
		expected     []int
		err          error
	}{
		{linesString: []string{"раз", "два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   false,
			invert:       false,
			fixed:        false,
			expected:     []int{1, 3},
			err:          nil,
		},
		{linesString: []string{"раз", "Два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   false,
			invert:       false,
			fixed:        false,
			expected:     []int{3},
			err:          nil,
		},
		{linesString: []string{"раз", "Два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   true,
			invert:       false,
			fixed:        false,
			expected:     []int{1, 3},
			err:          nil,
		},
		{linesString: []string{"раз", "Два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   true,
			invert:       true,
			fixed:        false,
			expected:     []int{0, 2, 4, 5},
			err:          nil,
		},
		{linesString: []string{"раз", "Два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   false,
			invert:       true,
			fixed:        false,
			expected:     []int{0, 1, 2, 4, 5},
			err:          nil,
		},
		{linesString: []string{"раз", "Два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   false,
			invert:       true,
			fixed:        true,
			expected:     []int{0, 1, 2, 4, 5},
			err:          nil,
		},
		{linesString: []string{"раз", "два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   false,
			invert:       false,
			fixed:        true,
			expected:     []int{1, 3},
			err:          nil,
		},
		{linesString: []string{"раз", "Два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   false,
			invert:       false,
			fixed:        true,
			expected:     []int{3},
			err:          nil,
		},
		{linesString: []string{"раз", "Два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   true,
			invert:       false,
			fixed:        true,
			expected:     []int{1, 3},
			err:          nil,
		},
		{linesString: []string{"раз", "Два", "три", "раз два три", "раз три четыре", "3 2 1"},
			targetString: "два",
			ignoreCase:   true,
			invert:       true,
			fixed:        true,
			expected:     []int{0, 2, 4, 5},
			err:          nil,
		},
	}

	for _, testcase := range testTable {
		result, err := searchNumberLinesByTargetString(testcase.linesString, testcase.targetString, testcase.ignoreCase, testcase.invert, testcase.fixed)
		assert.Equal(t, testcase.expected, result)
		assert.Equal(t, testcase.err, err)
	}
}

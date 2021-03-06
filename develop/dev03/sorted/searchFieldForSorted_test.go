package sorted

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchFieldForSorted(t *testing.T) {
	testTable := []struct {
		inputStrings string
		column       int
		expected     string
		err          error
	}{
		{
			inputStrings: "",
			column:       2,
			expected:     "",
			err:          errors.New("index column out of range0"),
		},
		{
			inputStrings: "Привет новый мир",
			column:       1,
			expected:     "новый",
			err:          nil,
		},
	}

	for _, testcase := range testTable {
		result, err := searchFieldFotSorted(testcase.inputStrings, testcase.column)
		assert.Equal(t, testcase.expected, result)
		assert.Equal(t, testcase.err, err)
	}
}

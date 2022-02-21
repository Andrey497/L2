package sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	testTable := []struct {
		inputStrings []string
		expected     []string
	}{
		{
			inputStrings: []string{"тест", "тест2"},
			expected:     []string{"тест", "тест2"},
		},
		{
			inputStrings: []string{"ямы", "ямы", "тест", "ямы"},
			expected:     []string{"ямы", "тест"},
		},
	}

	for _, testcase := range testTable {
		result := deleteDuplicates(testcase.inputStrings)
		assert.Equal(t, testcase.expected, result)
	}
}

package sorted

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedByNumber(t *testing.T) {
	testTable := []struct {
		lines    []string
		column   int
		reverse  bool
		expected []string
		err      error
	}{
		{
			lines:    []string{""},
			column:   2,
			reverse:  false,
			expected: []string{},
			err:      errors.New("index column out of range0"),
		},
		{
			lines:    []string{"а 1 я", "а 3 с", "б 2 ф", "к 3 с"},
			column:   1,
			reverse:  false,
			expected: []string{"а 1 я", "б 2 ф", "а 3 с", "к 3 с"},
			err:      nil,
		},
		{
			lines:    []string{"а 1 я", "а 3 с", "б 2 ф", "к 3 с"},
			column:   1,
			reverse:  true,
			expected: []string{"а 3 с", "к 3 с", "б 2 ф", "а 1 я"},
			err:      nil,
		},
	}

	for _, testcase := range testTable {
		result, err := sortedByNumber(testcase.column, testcase.lines, testcase.reverse)
		assert.Equal(t, testcase.expected, result)
		assert.Equal(t, testcase.err, err)
	}
}

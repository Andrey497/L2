package sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedString(t *testing.T) {
	testTable := []struct {
		lines       []string
		reverse     bool
		spaceIgnore bool
		expected    []string
	}{
		{
			lines:       []string{""},
			reverse:     false,
			spaceIgnore: false,
			expected:    []string{""},
		},
		{
			lines:       []string{" а2 1 я", "а1 3 с", "б 2 ф", "к 3 с"},
			reverse:     false,
			spaceIgnore: true,
			expected:    []string{"а1 3 с", "а2 1 я", "б 2 ф", "к 3 с"},
		},
		{
			lines:       []string{"а2 1 я", "а1 3 с", "б 2 ф", "к 3 с"},
			reverse:     true,
			spaceIgnore: false,
			expected:    []string{"к 3 с", "б 2 ф", "а2 1 я", "а1 3 с"},
		},
	}

	for _, testcase := range testTable {
		result := sortedString(testcase.lines, testcase.reverse, testcase.spaceIgnore)
		assert.Equal(t, testcase.expected, result)

	}
}

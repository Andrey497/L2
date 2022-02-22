package cut

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCutString(t *testing.T) {
	testTable := []struct {
		textLines []string
		fields    []int
		delimiter string
		expected  []string
	}{
		{
			textLines: []string{"1раз\t1два\t1три\t1четыре\t1пять", "2раз\t2два\t2три\t2четыре\t2пять", "3раз\t3два\t3три\t3четыре\t3пять", "4раз\t4два\t4три\t4четыре\t4пять"},
			fields:    []int{0, 2},
			delimiter: "\t",
			expected:  []string{"1раз\t1три", "2раз\t2три", "3раз\t3три", "4раз\t4три"},
		},
		{
			textLines: []string{"1раз,1два,1три,1четыре,1пять", "2раз,2два,2три,2четыре,2пять", "3раз,3два,3три,3четыре,3пять", "4раз,4два,4три,4четыре,4пять"},
			fields:    []int{0, 1},
			delimiter: ",",
			expected:  []string{"1раз,1два", "2раз,2два", "3раз,3два", "4раз,4два"},
		},
		{
			textLines: []string{"1раз,1два,1три,1четыре,1пять", "2раз 2два 2три 2четыре 2пять", "3раз,3два,3три,3четыре,3пять", "4раз,4два,4три,4четыре,4пять"},
			fields:    []int{0, 1},
			delimiter: ",",
			expected:  []string{"1раз,1два", "2раз 2два 2три 2четыре 2пять", "3раз,3два", "4раз,4два"},
		},
		{
			textLines: []string{"1раз,1два,1три,1четыре,1пять", "2раз 2два 2три 2четыре 2пять", "3раз,3два,3три,3четыре,3пять", "4раз,4два,4три,4четыре,4пять"},
			fields:    []int{},
			delimiter: ",",
			expected:  []string{"", "2раз 2два 2три 2четыре 2пять", "", ""},
		},
	}

	for _, testcase := range testTable {
		result := cutStrings(testcase.textLines, testcase.fields, testcase.delimiter)
		assert.Equal(t, testcase.expected, result)

	}
}

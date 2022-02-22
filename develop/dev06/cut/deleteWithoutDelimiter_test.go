package cut

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteWithoutDelimiter(t *testing.T) {
	testTable := []struct {
		textLines []string
		delimiter string
		expected  []string
	}{
		{
			textLines: []string{"1раз\t1два\t1три\t1четыре\t1пять", "2раз\t2два\t2три\t2четыре\t2пять", "3раз\t3два\t3три\t3четыре\t3пять", "4раз\t4два\t4три\t4четыре\t4пять"},
			delimiter: "\t",
			expected:  []string{"1раз\t1два\t1три\t1четыре\t1пять", "2раз\t2два\t2три\t2четыре\t2пять", "3раз\t3два\t3три\t3четыре\t3пять", "4раз\t4два\t4три\t4четыре\t4пять"},
		},
		{
			textLines: []string{"1раз,1два,1три,1четыре,1пять", "2раз,2два,2три,2четыре,2пять", "3раз,3два,3три,3четыре,3пять", "4раз,4два,4три,4четыре,4пять"},
			delimiter: ",",
			expected:  []string{"1раз,1два,1три,1четыре,1пять", "2раз,2два,2три,2четыре,2пять", "3раз,3два,3три,3четыре,3пять", "4раз,4два,4три,4четыре,4пять"},
		},
		{
			textLines: []string{"1раз,1два,1три,1четыре,1пять", "2раз 2два 2три 2четыре 2пять", "3раз,3два,3три,3четыре,3пять", "4раз,4два,4три,4четыре,4пять"},
			delimiter: ",",
			expected:  []string{"1раз,1два,1три,1четыре,1пять", "3раз,3два,3три,3четыре,3пять", "4раз,4два,4три,4четыре,4пять"},
		},
	}

	for _, testcase := range testTable {
		result := deleteWithoutDelimiter(testcase.textLines, testcase.delimiter)
		assert.Equal(t, testcase.expected, result)

	}
}

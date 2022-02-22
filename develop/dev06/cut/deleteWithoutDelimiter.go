package cut

import (
	"strings"
)

func deleteWithoutDelimiter(textLines []string, delimiter string) []string {
	result := []string{}
	for _, line := range textLines {
		if strings.Contains(line, delimiter) {
			result = append(result, line)
		}
	}
	return result
}

package greep

import (
	"regexp"
	"strings"
)

func searchNumberLinesByTargetString(linesText []string, targetString string, ignoreCase bool, invert bool, fixed bool) ([]int, error) {
	numberLinesResult := []int{}
	if ignoreCase {
		targetString = strings.ToLower(targetString)
	}

	for i, line := range linesText {
		if ignoreCase {
			line = strings.ToLower(line)
		}
		if !fixed { //regular
			matched, err := regexp.MatchString(targetString, line)
			if err != nil {
				return nil, err
			}
			if invert {
				matched = !matched
			}
			if matched {
				numberLinesResult = append(numberLinesResult, i)
			}
		} else { //contains
			contain := strings.Contains(line, targetString)
			if invert {
				contain = !contain
			}
			if contain {
				numberLinesResult = append(numberLinesResult, i)
			}
		}
	}
	return numberLinesResult, nil
}

package sorted

func checkSort(lines []string, sortLines []string) bool {
	if len(lines) != len(sortLines) {
		return false
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] != sortLines[i] {
			return false
		}
	}
	return true
}

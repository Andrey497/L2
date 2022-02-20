package sorted

func deleteDuplicates(lines []string) []string {
	result := []string{}
	setString := map[string]struct{}{}
	for _, val := range lines {
		if _, contains := setString[val]; !contains {
			setString[val] = struct{}{}
			result = append(result, val)
		}
	}
	return result
}

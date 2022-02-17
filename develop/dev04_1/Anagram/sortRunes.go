package Anagram

type SortRunes []rune

func CreateSortRunes(s string) SortRunes {
	var runes SortRunes
	runes = []rune(s)
	return runes
}

func (s *SortRunes) Len() int {
	return len(*s)
}

func (s *SortRunes) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SortRunes) Less(i int, j int) bool {
	return (*s)[i] < (*s)[j]
}

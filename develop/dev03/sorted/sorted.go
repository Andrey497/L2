package sorted

import (
	"errors"
	"flag"
	"fmt"
)

type sorted struct {
	textLines            []string
	sortLines            []string
	columnSort           int
	sortedByNumber       bool
	reverse              bool
	unique               bool
	sortedByMonth        bool
	spaceIgnore          bool
	flagCheckSort        bool
	sortedByNumberSuffix bool
	checkSort            string
}

func InitSort(lines []string) (*sorted, error) {
	if len(lines) == 0 {
		return nil, errors.New("empty text")
	}
	return &sorted{lines, []string{}, -1,
			false, false, false, false, false, false, false, ""},
		nil
}

func (s *sorted) GetResult() []string {
	return s.sortLines
}
func (s *sorted) GetIsSort() string {
	return s.checkSort
}

func (s *sorted) Start() error {
	fmt.Println("Start")
	var err error
	s.sortLines = s.textLines
	if err != nil {
		return err
	}
	if s.unique {
		s.sortLines = deleteDuplicates(s.sortLines)
	}
	if s.sortedByNumber {
		s.sortLines, err = sortedByNumber(s.columnSort, s.sortLines, s.reverse)
		if err != nil {
			return err
		}
	} else if s.sortedByMonth {
		s.sortLines, err = sortedByMonth(s.columnSort, s.sortLines, s.reverse)
		if err != nil {
			return err
		}
	} else if s.sortedByNumberSuffix {

	} else if s.columnSort != 0 {
		s.sortLines, err = sortedStringByColumn(s.columnSort, s.sortLines, s.reverse, s.spaceIgnore)
		if err != nil {
			return err
		}
	} else {
		s.sortLines = sortedString(s.sortLines, s.reverse, s.spaceIgnore)
	}
	if s.flagCheckSort {
		if checkSort(s.textLines, s.sortLines) {
			s.checkSort = "отсортированно"
		} else {
			s.checkSort = "неотсортированно"
		}
	}

	return nil
}

func (s *sorted) SwitchFlags() error {

	k := flag.Int("k", 0, "column")
	n := flag.Bool("n", false, "sortByNumber")
	r := flag.Bool("r", false, "reverse")
	u := flag.Bool("u", false, "Unique")
	m := flag.Bool("m", false, "sortByMonth")
	b := flag.Bool("b", false, "spaceIgnore")
	c := flag.Bool("c", false, "checkSort")
	h := flag.Bool("h", false, "sortedByNumberSuffix")
	flag.Parse()
	if (*n == true && (*m == true || *h == true)) ||
		(*m == true && (*n == true || *h == true)) ||
		(*h == true && (*n == true || *m == true)) {
		return errors.New("multiple sorting parameters")
	}
	fmt.Println(*n)
	s.columnSort = *k
	s.sortedByNumber = *n
	s.reverse = *r
	s.unique = *u
	s.sortedByMonth = *m
	s.spaceIgnore = *b
	s.flagCheckSort = *c
	s.sortedByNumberSuffix = *h
	if *k < 0 {
		return errors.New("column <1")
	}
	if *k != 0 {
		s.columnSort = *k - 1
	}
	return nil
}

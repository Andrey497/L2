package sorted

import (
	"errors"
	"flag"
	"fmt"
)

type Sorted struct {
	textLines            []string
	columnSort           int
	sortedByNumber       bool
	reverse              bool
	Unique               bool
	sortedByMonth        bool
	spaceIgnore          bool
	flagCheckSort        bool
	sortedByNumberSuffix bool
	checkSort            string
}

func InitSort(lines []string, columnSort int) (*Sorted, error) {
	if len(lines) == 0 {
		return nil, errors.New("empty text")
	}
	return &Sorted{lines, -1,
			false, false, false, false, false, false, false, ""},
		nil
}

func (s *Sorted) Start() error {
	err := s.switchFlags()
	if err != nil {
		return err
	}
	if s.flagCheckSort {

	}
	if s.Unique {

	}

	if s.sortedByNumber {

	} else if s.sortedByMonth {

	} else if s.sortedByNumberSuffix {

	} else {
		fmt.Println("")
	}

	return nil
}

func (s *Sorted) switchFlags() error {
	k := flag.Int("k", 0, "column")
	n := flag.Bool("n", false, "sortByNumber")
	r := flag.Bool("r", false, "reverse")
	u := flag.Bool("u", false, "Unique")
	m := flag.Bool("m", false, "sortByMonth")
	b := flag.Bool("b", false, "spaceIgnore")
	c := flag.Bool("m", false, "checkSort")
	h := flag.Bool("h", false, "sortedByNumberSuffix")

	flag.Parse()

	if (*n == true && (*m == true || *h == true)) ||
		(*m == true && (*n == true || *h == true)) ||
		(*h == true && (*n == true || *m == true)) {
		return errors.New("multiple sorting parameters")
	}
	s.columnSort = *k
	s.sortedByNumber = *n
	s.reverse = *r
	s.Unique = *u
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

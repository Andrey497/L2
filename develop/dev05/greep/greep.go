package greep

import (
	"errors"
	"flag"
	"fmt"
)

type greep struct {
	textLines      []string
	targetString   string
	resultLines    []string
	after          int
	before         int
	count          int
	countFlag      bool
	ignoreCaseFlag bool
	invertFlag     bool
	fixedFlag      bool
	lineNumFlag    bool
}

func InitGreep(lines []string, targetString string) (*greep, error) {
	if len(lines) == 0 {
		return nil, errors.New("empty text")
	}
	if targetString == "" {
		return nil, errors.New("empty target string")
	}
	return &greep{lines, targetString, []string{}, 0,
			0, 0, false, false, false, false, false},
		nil
}
func (g *greep) Start() error {
	err := g.switchFlags()
	if err != nil {
		return err
	}
	numbersLinesByTarget, err := searchNumberLinesByTargetString(g.textLines, g.targetString, g.ignoreCaseFlag, g.invertFlag, g.fixedFlag)
	if err != nil {
		return err
	}
	g.count = len(numbersLinesByTarget)
	g.resultLines = greepLinesSearched(g.textLines, numbersLinesByTarget, g.after, g.before, g.lineNumFlag)
	return nil
}

func (g *greep) GetResult() {
	if g.countFlag {
		fmt.Println(fmt.Sprintf("Count:%d", g.count))
	} else {
		for _, val := range g.resultLines {
			fmt.Println(val)
		}
	}
}

func (g *greep) switchFlags() error {
	A := flag.Int("A", 0, "after")
	B := flag.Int("B", 0, "before")
	C := flag.Int("C", 0, "context")
	c := flag.Bool("c", false, "count")
	i := flag.Bool("i", false, "ignore-case")
	v := flag.Bool("v", false, "invert")
	F := flag.Bool("F", false, "fixed")
	n := flag.Bool("n", false, "line num")
	flag.Parse()
	g.after = *A
	g.before = *B
	g.countFlag = *c
	g.ignoreCaseFlag = *i
	g.invertFlag = *v
	g.fixedFlag = *F
	g.lineNumFlag = *n
	if *C < 0 || *B != 0 || *C != 0 {
		return errors.New("context<0 or before<0 or after<0")
	}
	if (*C != 0) && (*B != 0 || *C != 0) {
		return errors.New("context!= 0 and (before!=0 or after!=0)")
	}
	if *C != 0 {
		g.after = *C
		g.before = *C
	}
	return nil
}

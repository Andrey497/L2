package cut

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
)

type cut struct {
	textLines            []string
	resultLines          []string
	fields               []int
	delimiter            string
	stringsOnlySeparated bool
}

func InitCut(textLines []string) (*cut, error) {
	if len(textLines) == 0 {
		return nil, errors.New("empty text")
	}
	return &cut{textLines, []string{}, []int{}, "", false},
		nil
}
func (c *cut) Start() error {
	err := c.switchFlags()
	c.resultLines = c.textLines
	if err != nil {
		return err
	}
	if c.stringsOnlySeparated {
		c.resultLines = deleteWithoutDelimiter(c.resultLines, c.delimiter)
	}
	c.resultLines = cutStrings(c.resultLines, c.fields, c.delimiter)
	return nil
}
func (c *cut) GetResult() []string {
	return c.resultLines
}
func (c *cut) PrintResult() {
	for _, line := range c.resultLines {
		fmt.Println(line)
	}
}

func (c *cut) switchFlags() error {
	//parse flag
	f := flag.String("f", "", "-f - \"fields\" - выбрать поля (колонки)")
	d := flag.String("d", "    ", "\"delimiter\" - использовать другой разделитель")
	s := flag.Bool("s", false, "\"separated\" - только строки с разделителем")
	flag.Parse()

	//convert string to list
	list := "[" + *f + "]"
	var intSlice []int
	err := json.Unmarshal([]byte(list), &intSlice)
	if err != nil {
		return err
	}
	c.delimiter = *d
	c.stringsOnlySeparated = *s
	c.fields = intSlice
	return nil
}

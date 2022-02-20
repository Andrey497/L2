package main

import "testing"

func TestUnPackageString(t *testing.T) {
	correctTestTable := []struct {
		inputString string
		expected    string
	}{
		{
			inputString: "a4bc2d5e",
			expected:    "aaaabccddddde",
		},
		{
			inputString: "abcd",
			expected:    "abcd",
		},
		{
			inputString: ``,
			expected:    "",
		},
		{
			inputString: `qwe\4\5`,
			expected:    "qwe45",
		},
		{
			inputString: `qwe\45`,
			expected:    `qwe44444`,
		},
		{
			inputString: `qwe\\5`,
			expected:    `qwe\\\\\`,
		},
	}
	uncorrectedTestValue := []string{"45"}
	for _, testcase := range correctTestTable {
		result, _ := UnPackageString(testcase.inputString)
		if result != testcase.expected {
			t.Errorf("Incorect unpackage function result on  correct value. Input value: %s, Output Value:%s", testcase.inputString, result)
		}
	}
	for _, testcase := range uncorrectedTestValue {
		result, err := UnPackageString(testcase)
		if err == nil {
			t.Errorf("Incorect unpackage function result on  Uncorrect value. Input value: %s, Output Value:%s", testcase, result)
		}
	}

}

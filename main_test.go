package main

import (
	"testing"
)

func TestDoStuff(t *testing.T) {
	csv1 := "COLUMN1,COLUMN2\ntest1,test2"
	csv2 := "COLUMN1|COLUMN2\ntest3|test4"
	s, err := doStuff(csv1, csv2)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if s != "Data1 Col1: test1 Data2 Col2: test4" {
		t.Errorf("String does not match expected: \t%v", s)
	}
}

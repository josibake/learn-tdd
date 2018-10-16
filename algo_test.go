package main

import (
	"reflect"
	"testing"
)

func TestCmdLineInputParsing(t *testing.T) {
	tables := []struct {
		i string
		o []string
	}{
		{"1+3*2", []string{"1", "+", "3", "*", "2"}},
		{"11+3*2", []string{"11", "+", "3", "*", "2"}},
		{"1+32*2", []string{"1", "+", "32", "*", "2"}},
		{"1+3*203", []string{"1", "+", "3", "*", "203"}},
		{"1.1+3*2", []string{"1.1", "+", "3", "*", "2"}},
		{"1.1+(3*2)", []string{"1.1", "+", "(", "3", "*", "2", ")"}},
		{"(1+1)*5+3", []string{"(", "1", "+", "1", ")", "*", "5", "+", "3"}},
	}

	for _, table := range tables {
		total := CmdLineInputParsing(table.i, ops)
		if reflect.DeepEqual(total, table.o) == false {
			t.Errorf("Parsing of %s was incorrect, got: %s, want: %s.", table.i, total, table.o)
		}
	}
}

func TestShuntingYardAlgorithm(t *testing.T) {
	tables := []struct {
		i []string
		o []string
	}{
		{[]string{"1", "+", "3", "*", "2"}, []string{"1", "3", "2", "*", "+"}},
		{[]string{"11", "+", "3", "*", "2"}, []string{"11", "3", "2", "*", "+"}},
		{[]string{"1", "+", "32", "*", "2"}, []string{"1", "32", "2", "*", "+"}},
		{[]string{"1", "+", "3", "*", "203"}, []string{"1", "3", "203", "*", "+"}},
		{[]string{"1.1", "+", "3", "*", "2"}, []string{"1.1", "3", "2", "*", "+"}},
	}

	for _, table := range tables {
		rpn := ShuntingYardAlgorithm(table.i, ops)
		if reflect.DeepEqual(rpn, table.o) == false {
			t.Errorf("Parsing of %s was incorrect, got: %s, want: %s.", table.i, rpn, table.o)
		}
	}
}

func TestComputeResult(t *testing.T) {
	tables := []struct {
		i []string
		o float64
	}{
		{[]string{"1", "3", "2", "*", "+"}, 7},
		{[]string{"11", "3", "2", "*", "+"}, 17},
		{[]string{"1", "32", "2", "*", "+"}, 65},
		{[]string{"1", "3", "203", "*", "+"}, 610},
		{[]string{"1.1", "3", "2", "*", "+"}, 7.1},
	}

	for _, table := range tables {
		result := ComputeResult(table.i, ops)
		if result != table.o {
			t.Errorf("Parsing of %s was incorrect, got: %f, want: %f.", table.i, result, table.o)
		}
	}
}
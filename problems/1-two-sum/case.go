package main

import "fmt"

type Input struct {
	Nums   []int
	Target int
}

type Output struct {
	Value []int
}

func (o Output) String() string {
	return fmt.Sprint(o.Value)
}

type Case struct {
	input  Input
	output Output
}

func GenerateCase() []Case {
	cases := []Case{
		{input: Input{Nums: []int{2, 7, 11, 15}, Target: 9}, output: Output{Value: []int{0, 1}}},
		{input: Input{Nums: []int{3, 2, 4}, Target: 6}, output: Output{Value: []int{1, 2}}},
	}
	return cases
}

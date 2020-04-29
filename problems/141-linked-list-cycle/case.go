package main

import "github.com/minidonut/leetcode-go/pkg/containers"

type Input = containers.ListNodeIntCyclic
type Output = containers.Boolean

type Case struct {
	input  Input
	output Output
}

func GenerateCase() []Case {
	node0 := Input{Val: 3, Next: nil}
	node1 := Input{Val: 2, Next: nil}
	node2 := Input{Val: 0, Next: nil}
	node3 := Input{Val: -4, Next: nil}

	node0.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node1

	cases := []Case{
		{input: node0, output: Output{Value: true}},
	}
	return cases
}

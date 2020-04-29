package containers

import (
	"strconv"
)

type ListNodeIntCyclic struct {
	Val  int
	Next *ListNodeIntCyclic
}

type ListNodeInt struct {
	Val  int
	Next *ListNodeInt
}

func (l *ListNodeInt) FromSlice(arr []int) ListNodeInt {
	head := ListNodeInt{
		Val:  arr[0],
		Next: nil,
	}

	cursor := &head

	for _, val := range arr[1:] {
		node := ListNodeInt{
			Val:  val,
			Next: nil,
		}

		cursor.Next = &node
		cursor = &node
	}

	return head
}

func (l ListNodeInt) String() string {
	s := ""
	cursor := &l
	for cursor != nil {
		s = s + "[" + strconv.Itoa(cursor.Val) + "]" + "-"
		cursor = cursor.Next
	}

	return s + "[nil]"
}

package containers

import "strconv"

type Integer struct {
	Value int
}

func (i Integer) String() string {
	return strconv.Itoa(i.Value)
}

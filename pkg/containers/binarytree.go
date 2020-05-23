package containers

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *

        0
   1          2
 3   4     5     6
7 8 9 10 11 12 13 14
*/

/**
 * 재밌는 아이큐테스트 아이디어
 *  1 -> 0, 8 -> 3, 12 -> 5, 16 -> 7, 10 -> ?
 */

func (t *TreeNode) fromSlice(arr []int) *TreeNode {
	root := TreeNode{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	table := map[int]*TreeNode{
		0: &root,
	}

	for i, v := range arr {
		if i == 0 {
			continue
		}
		// parentIndex := (i - 1)/2
		if p, ok := table[(i-1)/2]; ok {
			node := TreeNode{
				Val:   v,
				Left:  nil,
				Right: nil,
			}
			if i%2 == 1 {
				p.Left = &node
			} else {
				p.Right = &node
			}
		}
	}

	return &root
}

func (t TreeNode) String() string {
	queue := []*TreeNode{&t}
	result := []int{}

	for {
		if len(queue) == 0 {
			break
		}
		var node *TreeNode
		node, queue = queue[0], queue[1:]
		if node == nil {
			result = append(result, 0)
			continue
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	return fmt.Sprint(result)
}

package main

func Solve(nums []int, target int) []int {
	table := map[int]int{}

	for i, v := range nums {
		if first, ok := table[v]; ok {
			return []int{first, i}
		} else {
			table[target-v] = i
		}
	}

	return nil
}

/*
func Solve(nums []int, target int) []int {
	for i, v := range nums {
		result := []int{i}
		rest := target - v
		for j := i + 1; j < len(nums); j++ {
			if rest == nums[j] {
				return append(result, j)
			}
		}
	}
	return []int{0, 0}
}
*/

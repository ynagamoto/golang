package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	new_nums := make([][]int, len(nums))
	for i, _ := range new_nums {
		new_nums[i] = []int{nums[i], i}
	}

	// 小さい順にソート
	sort.Slice(new_nums, func(i, j int) bool { return new_nums[i][0] < new_nums[j][0]})

	result := []int{-1, -1}
	for i, v := range new_nums {
		rest := target - v[0]
		// rest以上 を二分探索
		index := sort.Search(len(new_nums), func(j int) bool { return i != j && new_nums[j][0] >= rest })
		if index != len(nums) { // 存在した
			if rest == new_nums[index][0] { // rest と一致
				result[0] = v[1]
				result[1] = new_nums[index][1] 
				break
			}
		}
	}

	if result[0] > result[1] {
		tmp := result[0]
		result[0] = result[1]
		result[1] = tmp
	}
	return result
}

func main() {
	// nums := []int{2, 7, 11, 15}
	nums := []int{15, 7, 11, 2}
	target := 9
	res := twoSum(nums, target)
	fmt.Printf("[%d, %d]\n", res[0], res[1])
}
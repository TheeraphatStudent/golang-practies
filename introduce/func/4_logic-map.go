package ilovemyjobPkg

import "fmt"

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, found := seen[complement]; found {
			return []int{j, i}
		}

		seen[num] = i
	}

	return nil
}

func LogicMap() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := TwoSum(nums, target)
	fmt.Println("Indices:", result)
	fmt.Println("Values:", nums[result[0]], "+", nums[result[1]])
}

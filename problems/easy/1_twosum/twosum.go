package twosum

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if idx, ok := idxMap[diff]; ok {
			return []int{idx, i}
		}
		idxMap[num] = i
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for i, v := range nums{
		a := target - v

		if prevIdx, ok := seen[a]; ok{
			return []int{prevIdx, i}
		}
		seen[v] = i
	}
	return []int{}
}

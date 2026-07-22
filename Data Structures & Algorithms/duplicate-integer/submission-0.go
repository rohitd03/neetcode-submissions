func hasDuplicate(nums []int) bool {
	num := make(map[int]bool)
	for  _, val := range nums{
		if _, ok := num[val]; ok {
			return true
		}
		num[val] = true
	}
	return false
}

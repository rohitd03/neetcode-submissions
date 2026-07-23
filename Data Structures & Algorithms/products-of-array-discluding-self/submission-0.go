func productExceptSelf(nums []int) []int {
	res := make([]int, 0, len(nums))
	for i, _ := range nums{
		prd := 1
		for j:= 0; j < len(nums); j++{
			if j!=i{
				prd *= nums[j]
			}
		}
		res = append(res, prd)
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	resMap := make(map[[26]int][]string)

	for _, w := range strs{
		var count [26]int
		for _, c := range w {
			count[c-'a']++
		}
		resMap[count] = append(resMap[count], w)
	}

	var res [][]string
	for _, group := range resMap {
		res = append(res, group)
	}
	return res
}

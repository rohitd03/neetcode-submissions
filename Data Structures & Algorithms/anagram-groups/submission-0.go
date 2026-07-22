func groupAnagrams(strs []string) [][]string {
	sortString := func(s string)string{
		r := []rune(s)

		sort.Slice(r, func(i, j int)bool{
			return r[i] < r[j]
		})

		return string(r)
	}

	groups := make(map[string][]string)

	for _, v := range strs{
		key := sortString(v)

		groups[key] = append(groups[key], v)
	}
	result := make([][]string,0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

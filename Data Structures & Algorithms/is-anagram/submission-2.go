func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
        return false
    }
	charsS := make(map[rune]int)
	charsT := make(map[rune]int)

	for i, v := range s {
		charsS[v]++
		charsT[rune(t[i])]++
	}
	for k, v := range charsS {
		if charsT[k] != v{
			return false
		}
	}
	return true
}

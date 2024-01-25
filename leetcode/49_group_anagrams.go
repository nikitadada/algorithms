package leetcode

type ArrayKey [26]byte

func strToArrayKey(str string) (key ArrayKey) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[ArrayKey][]string)

	for _, v := range strs {
		key := strToArrayKey(v)
		groups[key] = append(groups[key], v)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

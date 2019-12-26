package _028

/*
	实现 strStr() 函数。

	给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

	示例 1:
	输入: haystack = "hello", needle = "ll"
	输出: 2

	示例 2:
	输入: haystack = "aaaaa", needle = "bba"
	输出: -1
*/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	position := make([]int, 0)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			position = append(position, i)
		}
	}

	for _, pos := range position {
		i := 1
		for ; i < len(needle); i++ {
			if haystack[pos+i] == needle[i] {
				continue
			} else {
				break
			}
		}
		if i == len(needle) {
			return pos
		}
	}
	return -1
}

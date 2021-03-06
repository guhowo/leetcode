package _028

import (
	"testing"
)

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
func Test_StrStr(t *testing.T) {
	for _, unit := range []struct {
		Haystack string
		Needle   string
		Expect   int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	} {
		if actual := strStr(unit.Haystack, unit.Needle); actual != unit.Expect {
			t.Errorf("input = %v, actual = %d", unit, actual)
		}
	}
}

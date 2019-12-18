package _91

/*
	一条包含字母 A-Z 的消息通过以下方式进行了编码：

	'A' -> 1
	'B' -> 2
	...
	'Z' -> 26
	给定一个只包含数字的非空字符串，请计算解码方法的总数。

	示例 1:

	输入: "12"
	输出: 2
	解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
	示例 2:

	输入: "226"
	输出: 3
	解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

	链接：https://leetcode-cn.com/problems/decode-ways
*/

/*
	思路：
	利用动态转移方程，dp[i]=dp[i-1]+dp[i-2]
	但是需要考虑下面情况：
	若s[i]=0,则s[i-1]=2或者1才符合，即s[i-1]+s[i]是唯一破译方法，所以dp[i]对于dp[i-2]没有增加破译方法，dp[i]=dp[i-2]
	若s[i-1]=1,则dp[i]=dp[i-1]+dp[i-2]
	若s[i-1]=2,if 1<s[i]<=6,则dp[i]=dp[i-1]+dp[i-2]

	最终解释：s[i-1]和s[i]分开译：dp[i]=dp[i-1]，
			一起译dp[i]=dp[i-2]；
			所以dp[i]=dp[i-1]+dp[i-2];
			每次判断的时候就看能不能分开翻译。
*/
func numDecodings(s string) int {
	states := make([]int, len(s))
	if s[0] == '0' {
		return 0
	} else {
		states[0] = 1
	}

	if len(s) >= 2 {
		if s[1] == '0' {
			if s[0] == '1' || s[0] == '2' {
				states[1] = 1
			} else {
				return 0
			}
		} else if s[0] == '1' {
			states[1] = 2
		} else if s[0] == '2' && (s[1] >= '1' && s[1] <= '6') {
			states[1] = 2
		} else {
			states[1] = 1
		}
	}
	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				states[i] = states[i-2]
			} else {
				return 0
			}
		} else {
			if s[i-1] == '1' {
				states[i] = states[i-1] + states[i-2]
			} else if s[i-1] == '2' {
				if s[i] >= '1' && s[i] <= '6' {
					states[i] = states[i-1] + states[i-2]
				} else {
					states[i] = states[i-1]
				}
			} else {
				states[i] = states[i-1]
			}
		}
	}

	return states[len(s)-1]
}

//for i := 2; i < len(states); i++ {
//if s[i] == '0' {
//if s[i-1] != '1' && s[i-1] != '2' {
//return 0
//}
//if states[i-1] == 1 {
//states[i] = 1
//} else {
//states[i] = states[i-1] - 1
//}
//} else if s[i-1:i+1] > "26" || s[i-1] == '0' {
//states[i] = states[i-1]
//} else if s[i-1:i+1] <= "26" && s[i-1:i+1] >= "11" {
//states[i] = states[i-1] + 1
//}
//}

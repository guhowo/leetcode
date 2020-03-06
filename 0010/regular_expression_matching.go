package _010

func isMatch(s string, p string) bool {
	matched := false
	match(0, 0, s, p)
	return matched
}

func match(i, j int, s, p string) bool {
	//全部对比结束
	if i == len(s) && j == len(p) {
		return true
	}

	if i == len(s) || j == len(p) {
		return false
	}

	if s[i] == p[j] {
		return match(i+1, j+1, s, p)
	} else {
		//匹配任意单个字母
		if p[j] == '.' {
			return match(i+1, j+1, s, p)
		}
		if p[j] == '*' {
			return match(i+1, j+1, s, p) || match(i, j+1, s, p)
		}
		return false
	}
}

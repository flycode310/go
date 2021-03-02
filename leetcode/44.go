package leetcode

func isMatch(s string, p string) bool {
	if p == "*" || (s == "" && p == "") {
		return true
	}
	if s == "" || p == "" {
		return false
	}
	sIndex, pIndex := 0,0
	sLen, pLen := len(s), len(p)

	match := true
	for ; sIndex<= sLen; sIndex++ {
		if pIndex >= pLen {
			break
		}
		if p[pIndex] == '*' {
			continue
		}
		if p[pIndex] == '?' {
			pIndex ++
			continue
		}
		if s[sIndex] == p[pIndex] {
			pIndex ++
			continue
		}
		match = false
		break
	}

	if sIndex < sLen {
		match = false
	}
	return match
}

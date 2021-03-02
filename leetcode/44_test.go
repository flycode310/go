package leetcode

import (
	"testing"
)

func TestIsMatch(t *testing.T)  {
	matchStr(t, "aa", "a", false)
	matchStr(t, "aa", "*", true)
	matchStr(t, "cb", "?a", false)
	matchStr(t, "adceb", "*a*b", true)
	matchStr(t, "acdcb", "a*c?b", false)
}

func matchStr(t *testing.T, s string, p string, match bool) {
	if isMatch(s, p) != match {
		t.Error(s + ":" + p)
	}
}
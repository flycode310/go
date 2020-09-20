package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	s2TMap := map[rune]rune{}
	t2SMap := map[rune]rune{}
	strLen := len(s)
	for i:=0; i<strLen; i++ {
		toChar, ok := s2TMap[rune(s[i])]
		if !ok {
			toChar = rune(t[i])
			if _, isConflict := t2SMap[toChar]; isConflict {
				return false
			}
			s2TMap[rune(s[i])] = toChar
			t2SMap[toChar] = rune(s[i])
			continue
		}
		if toChar != rune(t[i]) {
			return false
		}
	}
	return true

}

func main()  {

	ret := isIsomorphic("egg", "add")
	fmt.Println(ret)
	ret = isIsomorphic("foo", "bar")
	fmt.Println(ret)
	ret = isIsomorphic("paper", "title")
	fmt.Println(ret)
	ret = isIsomorphic("ab", "aa")
	fmt.Println(ret)
}

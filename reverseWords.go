package main

import "fmt"

func reverseWords(s string) string {
	retStr := ""
	low := 0
	high := 0
	for i:=0; i<len(s); i++ {
		if i == 0 || s[i-1] == ' ' {
			low = i
			if i > 0 {
				retStr += " "
			}
		}
		if (i == len(s)-1) || s[i+1] == ' ' {
			high = i
			for j:=high; j>=low; j-- {
				retStr += string(s[j])
			}
		}

	}
	return retStr
}

func main() {
	str := "Let's take LeetCode contest"
	//str = "a good   example"
	str = reverseWords(str)
	fmt.Printf("%q", str)
}

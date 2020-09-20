package main

import "fmt"

func generateParenthesis(n int) []string {
	ret := &[]string{}
	generate(ret, "(", n-1, n)
	return *ret
}

func generate(ret *[]string, str string, leftPCnt int, rightPCnt int) {
	if leftPCnt > rightPCnt {
		return
	}
	if leftPCnt == 0 && rightPCnt == 0 {
		*ret = append(*ret, str)
		return
	}
	if leftPCnt > 0 {
		generate(ret, str+"(", leftPCnt-1, rightPCnt)
	}
	generate(ret, str+")", leftPCnt, rightPCnt-1)
}

func main()  {
	ret := generateParenthesis(1)
	fmt.Printf("%q",ret)
}
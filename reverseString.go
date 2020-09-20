package main

func reverseString(s []byte)  {
	low := 0
	high := len(s) - 1
	tmp := byte(0)
	for low < high {
		tmp = s[low]
		s[low] = s[high]
		s[high] = tmp
		low ++
		high --
	}
}

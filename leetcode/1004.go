package leetcode

func longestOnes(A []int, K int) int {
	maxNums, curNums := 0, 0
	for index, value := range A {
		if value == 1 {
			curNums ++
			if curNums > maxNums {
				maxNums = curNums
			}
		} else {
			sIndex := index
			tmpK := K
			for ; sIndex < len(A); sIndex ++ {
				if A[sIndex] == 0 {
					tmpK --
					if tmpK < 0 {
						break
					}
				}
				curNums ++
			}
			if curNums > maxNums {
				maxNums = curNums
			}
			curNums = 0;
		}
	}
	return maxNums
}

func longestOneA(A []int, K int) int {
	maxNums, curNums := 0, 0
	for index, value := range A {
		if value == 1 {
			curNums ++
			if curNums > maxNums {
				maxNums = curNums
			}
		} else {
			sIndex := index
			tmpK := K
			for ; sIndex < len(A); sIndex ++ {
				if A[sIndex] == 0 {
					tmpK --
					if tmpK < 0 {
						break
					}
				}
				curNums ++
			}
			if curNums > maxNums {
				maxNums = curNums
			}
			curNums = 0;
		}
	}
	return maxNums
}

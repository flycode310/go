package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	var matrix [][]int
	var row []int
	for i:=0; i<=len2; i++ {
		row = append(row, i)
	}
	matrix = append(matrix, row)
	for i:=1; i<=len1; i++ {
		row = []int{}
		for j:=0; j<=len2; j++ {
			row = append(row, 0)
		}
		matrix = append(matrix, row)
	}
	for i:=0; i<=len1; i++ {
		matrix[i][0] = i
	}

	minDistance := 0
	delDistance :=0
	addDistance := 0
	modDistance := 0
	for i:=1; i<=len1; i++ {
		for j:=1; j<=len2; j++ {
			delDistance = matrix[i-1][j] + 1
			addDistance = matrix[i][j-1] + 1
			modDistance = matrix[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				modDistance ++
			}
			matrix[i][j] = min(delDistance, addDistance, modDistance)
		}
	}
	minDistance = matrix[len1][len2]

	return minDistance
}

func min(values ...int) int {
	minNum := 0xFFFFFFFF
	for _, value := range values {
		if value < minNum {
			minNum = value
		}
	}
	return minNum
}

func main()  {
	str1 := "intention"
	str2 := "execution"
	distance := minDistance(str1, str2)
	fmt.Println(distance)

}

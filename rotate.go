package main

import "fmt"

func rotate(matrix [][]int)  {
	n := len(matrix)
	ringStartX := 0
	ringStartY := 0
	ringTotalNum := n / 2
	for ringNo:=0; ringNo<ringTotalNum; ringNo++ {
		ringStartX += 1
		ringStartY += 1
		rightY := n -1 - ringStartX -1
		for startY := ringStartY; startY<=rightY; startY++ {
			startX := ringStartX
			fromX := startX
			fromY := startY
			for {
				toX := fromY
				toY := n - 1 - fromX
				if (toX == startX) && (toY == startY) {
					break
				}
				swapPosition := [4]int{startX, startY, toX, toY}
				swap(matrix, swapPosition)
				fromX = toX
				fromY = toY
			}
		}

	}
}

func swap(matrix [][]int, position [4]int)  {
	startX := position[0]
	startY := position[1]
	endX := position[2]
	endY := position[3]
	matrix[endX][endY] += matrix[startX][startY]
	matrix[startX][startY] = matrix[endX][endY] - matrix[startX][startY]
	matrix[endX][endY] -= matrix[startX][startY]
}

func main() {
	//matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	matrix := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	rotate(matrix)
	fmt.Printf("%v\n", matrix)
}

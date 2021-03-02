package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	rowNum, columnNum := len(matrix), len(matrix[0])
	rowStep, columnStep := 0, 1
	indexR, indexC := 0, 0
	ant := make([]int, rowNum * columnNum)
	rHighIndex, rLowIndex, cHighIndex, cLowIndex := rowNum, -1, columnNum, -1
	for i:=0; i<rowNum * columnNum; i++ {
		ant[i] = matrix[indexR][indexC]
		nextR := indexR + rowStep
		nextC := indexC + columnStep
		if rowStep == 0 && columnStep == 1 && nextC >= cHighIndex {
			rowStep = 1
			columnStep = 0
			rLowIndex ++
		} else if rowStep == 1 && columnStep == 0 && nextR >= rHighIndex {
			rowStep = 0
			columnStep = -1
			cHighIndex --
		} else if rowStep == 0 && columnStep == -1 && nextC <= cLowIndex {
			rowStep = -1
			columnStep = 0
			rHighIndex --
		} else if rowStep == -1 && columnStep == 0 && nextR <= rLowIndex {
			rowStep = 0
			columnStep = 1
			cLowIndex ++
		}
		indexR += rowStep
		indexC += columnStep
	}
	return ant
}

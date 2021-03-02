package leetcode

func findShortestSubArray(nums []int) int {
	mapGrade := map[int]int{}
	mapBegin := map[int]int{}
	mapEnd   := map[int]int{}
	grade := 1
	for index, value := range nums {
		if _, ok := mapGrade[value]; !ok {
			mapGrade[value] = 1
			mapBegin[value] = index
			mapEnd[value] = index
		} else {
			mapGrade[value] ++
			mapEnd[value] = index
		}

		if mapGrade[value] > grade {
			grade = mapGrade[value]
		}
	}

	ant := len(nums)
	for key, value := range mapGrade {
		if grade == value {
			tmpAnt := mapEnd[key] - mapBegin[key] + 1
			if tmpAnt < ant {
				ant = tmpAnt
			}
		}
	}
	return ant
}

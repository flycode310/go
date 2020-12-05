package leetcode

func numJewelsInStones(J string, S string) int {
	lSet := make(map[byte]byte)
	byteJ := []byte(J)
	for _,v := range byteJ {
		lSet[v] = 1
	}
	byteS := []byte(S)
	retNum := 0
	for _,v := range byteS {
		if _, ok := lSet[v]; ok {
			retNum ++
		}
	}
	return retNum
}

func NumJewelsInStones() {

}

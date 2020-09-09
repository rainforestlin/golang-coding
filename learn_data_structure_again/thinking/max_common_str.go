package thinking

func getCommon(a, b string) string {
	aArray := []byte(a)
	bArray := []byte(b)
	row := len(bArray)
	column := len(aArray)

	var commonArray = make([][]int, row+1)

	for i := 0; i < len(commonArray); i++ {
		commonArray[i] = make([]int, column+1)
	}

	for i := 1; i < len(bArray); i++ {
		for j := 1; j < len(aArray); j++ {
			if aArray[j-1] == bArray[i-1] {
				commonArray[i][j] = commonArray[i-1][j-1] + 1
			}
		}

	}

	var max int
	var index int
	for i := 0; i < len(aArray); i++ {
		for j := 0; j < len(bArray); j++ {
			if commonArray[j][i] > max {
				max = commonArray[j][i]
				index = i
			}
		}
	}

	var commonStr []byte
	for i := index - max; i < index; i++ {
		commonStr = append(commonStr, aArray[i])
	}
	return string(commonStr[:])

}

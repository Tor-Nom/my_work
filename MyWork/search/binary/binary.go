package binary

func Search(searchNum int, leftIndex, rightIndex int, searchArr []int) int {
	if leftIndex > rightIndex {
		return -1
	}

	middleIndex := (leftIndex + rightIndex) >> 1

	if searchNum < searchArr[middleIndex] {
		return Search(searchNum, leftIndex, middleIndex-1, searchArr)
	} else if searchNum > searchArr[middleIndex] {
		return Search(searchNum, middleIndex+1, rightIndex, searchArr)
	}
	return middleIndex

}

func SearchFor(searchNum int, leftIndex, rightIndex int, searchArr []int) int {
	for {
		middleIndex := (leftIndex + rightIndex) >> 1
		if searchNum < searchArr[middleIndex] {
			rightIndex = middleIndex - 1
		} else if searchNum > searchArr[middleIndex] {
			leftIndex = middleIndex + 1
		} else {
			return middleIndex
		}
	}
}

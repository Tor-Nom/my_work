package merge

import "fmt"

type MyMerge struct {
	SortArr []int
	Length  int
}

func NewMyMerge(sortArr []int) (*MyMerge, bool) {
	fmt.Println("Before MergeSort => ", sortArr)
	arrLen := len(sortArr)
	if arrLen == 0 || arrLen == 1 {
		return nil, false
	}
	return &MyMerge{
		SortArr: sortArr,
		Length:  arrLen,
	}, true
}

func (mm *MyMerge) Sort() {
	mm.SortArr = Divide(mm.SortArr)
}

func (mm *MyMerge) Print() {
	fmt.Println("MergeSort => ", mm.SortArr)
}

// 4,3,5,6, 2,7,9,1
func Divide(sortArr []int) []int {
	length := len(sortArr)
	if length == 1 {
		return sortArr
	}

	middle := length >> 1

	leftArr := sortArr[:middle]
	rightArr := sortArr[middle:]

	return Merge(Divide(leftArr), Divide(rightArr))
}

//3 4 ,5 6,2 7, 1 9
func Merge(leftArr, rightArr []int) []int {
	fmt.Println("leftArr, rightArr => ", leftArr, rightArr)
	mergeArr := make([]int, 0)

	for len(leftArr) != 0 && len(rightArr) != 0 {
		if leftArr[0] >= rightArr[0] {
			mergeArr = append(mergeArr, rightArr[0])
			rightArr = rightArr[1:]
		} else {
			mergeArr = append(mergeArr, leftArr[0])
			leftArr = leftArr[1:]
		}
	}

	if len(rightArr) != 0 {
		mergeArr = append(mergeArr, rightArr...)
	}
	if len(leftArr) != 0 {
		mergeArr = append(mergeArr, leftArr...)
	}
	return mergeArr
}

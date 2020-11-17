package _select

import (
	"fmt"
	"hello/sort/common"
)

type MySelect struct {
	SortArr []int
	Length  int
	Lib     common.Lib
}

func NewMySelect(sortArr []int) (*MySelect, bool) {
	arrLen := len(sortArr)
	if arrLen == 0 || arrLen == 1 {
		return nil, false
	}
	return &MySelect{
		SortArr: sortArr,
		Length:  arrLen,
	}, true
}
func (ms *MySelect) Sort() {
	for i := 0; i < ms.Length-1; i++ {
		maxIndex := 0
		for j := 0; j < ms.Length-i; j++ {
			//if ms.SortArr[j] >= ms.SortArr[maxIndex] {
			if ms.Lib.CompareStable(ms.SortArr[j], ms.SortArr[maxIndex]) {
				maxIndex = j
			}
		}

		ms.SortArr[ms.Length-1-i], ms.SortArr[maxIndex] = ms.SortArr[maxIndex], ms.SortArr[ms.Length-1-i]
		fmt.Println("第v%次排序完：", i, ms.SortArr)
	}
}

func (ms *MySelect) Print() {
	fmt.Println("SelectSort =>", ms.SortArr)
}

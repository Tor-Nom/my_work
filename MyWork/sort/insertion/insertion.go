package insertion

import (
	"fmt"
)

type MyInsertion struct {
	SortArr []int
	Length  int
}

func NewMyInsertion(sortArr []int) (*MyInsertion, bool) {
	arrLen := len(sortArr)
	if arrLen == 0 || arrLen == 1 {
		return nil, false
	}
	return &MyInsertion{
		SortArr: sortArr,
		Length:  arrLen,
	}, true
}
func (mi *MyInsertion) Sort() {
	for i := 1; i < mi.Length; i++ {
		for j := i; j > 0; j-- {
			if mi.SortArr[j-1] > mi.SortArr[j] {
				mi.SortArr[j-1], mi.SortArr[j] = mi.SortArr[j], mi.SortArr[j-1]
			} else {
				break
			}
		}
		fmt.Println("第v%次排序完：", i, mi.SortArr)
	}
}

func (mi *MyInsertion) Print() {
	fmt.Println("InsertionSort => ", mi.SortArr)
}
func (mi *MyInsertion) SortOpt() {
	for i := 1; i < mi.Length; i++ {
		tmp, endIndex := mi.SortArr[i], i

		for j := i; j > 0; j-- {
			if mi.SortArr[j-1] > tmp { //与前面排好序的集合比较，找到一个比当前tmp小的值
				mi.SortArr[j], endIndex = mi.SortArr[j-1], j-1
			}
		}
		mi.SortArr[endIndex] = tmp
		fmt.Println("第v%次排序完：", i, mi.SortArr)
	}
}

//func (mi *MyInsertion) SortOptWithBinarySearch() {
//	for i := 1; i < mi.Length; i++ {
//		tmp, endIndex := mi.SortArr[i], i
//
//		index := binary.SearchFor(tmp, 0, i, mi.SortArr)
//		for j := i; j > index; j-- {
//			mi.SortArr[j], endIndex = mi.SortArr[j-1], j-1
//		}
//		mi.SortArr[endIndex] = tmp
//		fmt.Println("第v%次排序完：", i, mi.SortArr)
//	}
//}
